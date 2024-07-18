package factory

import (
	"fmt"
	"golang-client/bpcontext"
	message "golang-client/message/protoData"
	"golang-client/modules/logger"

	"log"
	"math"
	"strconv"
	"strings"
)

func ExtractNodeType(nodeId uint64) uint64 {
	return nodeId / uint64(math.Pow10(8))
}
func IsDataNode(nodeId uint64) bool {
	return nodeId/uint64(math.Pow10(9)) == 1
}
func IsFunctionNode(nodeId uint64) bool {
	return nodeId/uint64(math.Pow10(9)) == 0
}
func deserializeFunctionNode(function_node *message.NodeData, ctx *bpcontext.DobitContext) *message.NodeData {
	log := logger.GetLogger().WithField("func", "deserializeFunctionNode")
	var functionPrompts string
	if function_node != nil && function_node.FunctionParam != nil {
		functionPrompts = *function_node.FunctionParam.FunctionPrompt
	}
	//log.Debugf("FunctionNode: %d Requested with Prompt %s", function_node.NodeId, functionPrompts)
	if function_node.NodeStructure == nil {
		log.Info("FunctionNode: No Child Node")
		return function_node
	}

	functionNodeData := &message.FunctionParams{
		FunctionPrompt: &functionPrompts,
	}
	nodeConnector := message.NodeConnector{}
	nodeConnector.InputNodes = make(map[int32]*message.NodeData)

	switch ExtractNodeType(function_node.NodeId) {
	case 1: // prompt input
		for index, node := range function_node.NodeStructure.InputNodes {
			if IsDataNode(node.NodeId) {
				//only for the service that with prompt inputs
				_, prompts := deserializeDataNode(node, ctx)
				//concatenate prompts
				functionPrompts = replaceIndexInPrompts(functionPrompts, index, prompts)
				//log.Print("FunctionPrompt:", functionPrompts)
				functionNodeData.FunctionPrompt = &functionPrompts
				continue
			}
			nodeConnector.InputNodes[index] = deserializeFunctionNode(node, ctx)
		}
	case 2: // Subservice Function (branch functions on root)
		for index, node := range function_node.NodeStructure.InputNodes {
			if IsDataNode(node.NodeId) {
				dataObject, _ := deserializeDataNode(node, ctx)
				functionNodeData.InputDataObj = dataObject
				continue
			}
			nodeConnector.InputNodes[index] = deserializeFunctionNode(node, ctx)
		}
	case 3: // TextInputMainServiceFunction
		functionNodeData.InputText = ctx.DataContext.InputText()
		for index, node := range function_node.NodeStructure.InputNodes {
			if IsDataNode(node.NodeId) {
				_, prompts := deserializeDataNode(node, ctx)
				//concatenate prompts
				functionPrompts = replaceIndexInPrompts(functionPrompts, index, prompts)
				//log.Print("FunctionPrompt:", functionPrompts)
				functionNodeData.FunctionPrompt = &functionPrompts
				//TODO <<<<<<<need Review

				//makelog.Debugf("InputText: %s With FunctionNode: %d", *functionNodeData.InputText, function_node.NodeId)
				continue
			}
			nodeConnector.InputNodes[index] = deserializeFunctionNode(node, ctx)
		}
	}
	//attach none string function_node structure into Input Data Object

	return &message.NodeData{
		NodeId:        function_node.NodeId,
		FunctionParam: functionNodeData,
		NodeStructure: &nodeConnector,
	}

}
func deserializeDataNode(data *message.NodeData, ctx *bpcontext.DobitContext) ([]byte, string) {
	//log := logger.GetLogger().WithField("func", "deserializeDataNode")

	switch ExtractNodeType(data.NodeId) {
	case 11: // General Data
		dataObj, dataString := deserializeData(data.NodeId, ctx)
		return dataObj, dataString
	case 12: // Abstract Data
		dataCtx := bpcontext.NewDataContext(ctx.DataContext)

		var inputString string
		for _, node := range data.NodeStructure.InputNodes {
			//Only accept the first protoData node
			if !IsDataNode(node.NodeId) {
				log.Fatalln("DataNode: Input Node for Data must be a protoData node")
			}
			_, inputString = deserializeDataNode(node, ctx)
			break
		}
		dataCtx.SetInputText(&inputString)
		dataObj, dataString := deserializeData(data.NodeId, ctx)
		//log.Debugf("DataString: %s", dataString)
		return dataObj, dataString

	default:
		dataObj, dataString := deserializeData(data.NodeId, ctx)
		//log.Debugf("DataString: %s", dataString)
		return dataObj, dataString
	}
}

// return ServiceFunction byte stream
func replaceIndexInPrompts(origPrompts string, index int32, returnPrompts string) string {
	//log := logger.GetLogger().WithField("func", "replaceIndexInPrompts")
	replacestr := "{" + strconv.Itoa(int(index)) + "}"
	//log.Debugf("index %v, orig: %v, replace: %v, return: %v", index, origPrompts, replacestr, returnPrompts)
	return strings.Replace(origPrompts, replacestr, returnPrompts, -1)
}
func DeserializeRootNode(data *message.NodeData, ctx *bpcontext.DobitContext) (*message.NodeData, error) {
	log := logger.GetLogger().WithField("func", "DeserializeRootNode")
	if data == nil {
		log.Errorf("RootNode: Root node is nil")
		return nil, fmt.Errorf("RootNode: Root node is nil")
	}
	log.Infof("Main Function Calling With Root Node: %d", data.NodeId)
	if !IsFunctionNode(data.NodeId) {
		log.Errorf("RootNode: Root node must be a function node")
		return nil, fmt.Errorf("RootNode: Root node must be a function node")
	}
	return deserializeFunctionNode(data, ctx), nil
}
