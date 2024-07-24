package dobit_client

import (
	"fmt"
	"golang-client/bpcontext"
	"golang-client/factory"
	"golang-client/implementation"
	"golang-client/py_comm_client"

	"golang-client/message/proto"
	"golang-client/modules/logger"
)

type Query struct {
	BPFunctionNodes *message.NodeData
}

func (q *Query) call(d *AgentEntity, ctx *bpcontext.AgentContext) {
	//q.Query_mu.Lock()
	//defer q.Query_mu.UnlockContext()
	log := logger.GetLogger().WithField("func", "call")
	log.Debugf("call: %v", q.BPFunctionNodes.NodeId)
	err := d.callMainFunction(q.BPFunctionNodes, ctx)
	if err != nil {
		log.Errorf("MainFunction call Suspended: %v", err)
		//if ctx.DataContext.Callback() != nil && ctx.DataContext.Callback().ResponseCancel != nil {
		//	ctx.DataContext.Callback().ResponseCancel()
		//}
	}
	//fmt.Printf("--------------End of Quring---------------")
}
func (d *AgentEntity) registerFunctionCall(functionNode *message.NodeData) {
	//call := d.Queries[functionNode.NodeId]
	//call.BPFunctionNodes =  onNode
	d.Queries[functionNode.NodeId] = Query{
		BPFunctionNodes: functionNode,
		//Query_mu:        sync.Mutex{},
	}
}

func (d *AgentEntity) callMainFunction(node *message.NodeData, ctx *bpcontext.AgentContext) error {
	log := logger.GetLogger().WithField("func", "callMainFunction")
	rqstNode, err := factory.DeserializeRootNode(node, ctx)
	if err != nil {
		log.Errorf("DeserializeRootNode error: %v", err)
		return err
	}
	response, err := py_comm_client.SendMainServiceRequest(rqstNode)
	if err != nil {
		log.Errorf("SendMainServiceRequest error: %v", err)
		return err
	}
	if response == nil {
		log.Errorf("response is nil with function_id: %d", node.NodeId)
		return fmt.Errorf("response is nil with function_id: %d", node.NodeId)

	}
	dataID := implementation.GetDataIDFromFunctionID(node.NodeId)
	datamgr := bpcontext.GetDataManager(dataID)
	if datamgr == nil {
		log.Errorf("GetDataManager error: %v, function_id: %v", dataID, node.NodeId)
		return fmt.Errorf("GetDataManager error: %v, function_id: %v", dataID, node.NodeId)
	}
	datamgr.SetServiceResponse(response.MessageId, response.ResData, d, ctx.DataContext)
	return nil
}
func (d *AgentEntity) callSubordinateFunction(function_id uint64, data_id uint64, data []byte, ctx bpcontext.QueryContextInterface) error {
	log := logger.GetLogger().WithField("func", "callSubordinateFunction")
	response, err := py_comm_client.SendSubordinateServiceRequest(function_id, data_id, data)
	if err != nil {
		log.Errorf("SendSubordinateServiceRequest error: %v", err)
		return err
	}
	if response == nil {
		log.Errorf("response is nil with function_id: %d", function_id)
		return nil
	}
	dataID := implementation.GetDataIDFromFunctionID(function_id)
	datamgr := bpcontext.GetDataManager(dataID)
	if datamgr == nil {
		log.Errorf("GetDataManager error: %v, function_id: %v", dataID, function_id)
		return fmt.Errorf("GetDataManager error: %v, function_id: %v", dataID, function_id)
	}
	datamgr.SetServiceResponse(response.MessageId, response.ResData, d, ctx)
	return nil
}
