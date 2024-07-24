package apm_util

import (
	apm "golang-client/message/proto"
	"golang-client/modules/logger"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func WriteGenericMethodAPM(rootNodeIds []uint64, filename string) {
	object := apm.ApmFile{
		Trees: []*apm.FileTree{},
	}
	for _, nodeId := range rootNodeIds {
		treeNode := writeSingleFunctionTreeNode(nodeId)
		object.Trees = append(object.Trees, treeNode)
	}

	out, err := proto.Marshal(&object)
	if err != nil {
		panic(err)
	}
	inputStreamToFile(out, filename)
}

func writeSingleFunctionTreeNode(funcId uint64) *apm.FileTree {

	var treeNode *apm.FileTree
	switch funcId {
	case 110000001: //generate events
		treeNode = getFileTree2GenerateEvents(funcId)
	default:
		logger.GetLogger().Printf("Invalid Node Id: %v", funcId)
	}
	return treeNode
}

func getFileTree2GenerateEvents(funcId uint64) *apm.FileTree {
	ActionDataNode := apm.NodeData{NodeId: 1110010000}
	ParsedActionDataNode := apm.NodeData{NodeId: 1140010000}

	rInputNodes := make(map[int32]*apm.NodeData)
	rInputNodes[1] = &ActionDataNode
	rInputNodes[2] = &ParsedActionDataNode
	rootNodeConnector := apm.NodeConnector{
		InputNodes: rInputNodes,
	}
	//testPrompt := "{1}"
	testPrompt := "Describe your prompt here to with placeholder {1}" +
		"and placeholder {2} will be replaced during runtime."
	getFunctions := apm.FunctionParams{
		FunctionPrompt: &testPrompt,
	}
	rootNode := apm.NodeData{
		NodeId:        funcId, //  100100800,generate events
		NodeStructure: &rootNodeConnector,
		FunctionParam: &getFunctions,
	}
	treeNode := apm.FileTree{
		RootNode: &rootNode,
	}
	return &treeNode
}

func inputStreamToFile(data []byte, filename string) {
	//open File test_example.go, if not exist create it
	f, createError := os.Create(filename)
	if createError != nil {
		log.Fatal(createError)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	_, err := f.Write(data)
	if err != nil {
		log.Fatal(err)
		return
	}

}
