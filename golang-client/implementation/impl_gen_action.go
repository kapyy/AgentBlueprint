package implementation

import (
	bpcontext "golang-client/bpcontext"
	proto "golang-client/message/proto"
	logger "golang-client/modules/logger"
	proto1 "google.golang.org/protobuf/proto"
)

func (m *ActionManager) Default(d bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me, this is where you read this data from, could be connected to a database or a service
	panic("implement me")
	/*
	   actions := &ActionList{}
	   actionList.Set(&proto.ActionList{
	   Actions:[]*proto.Action{
	   {
	   ActionDescription: "go for a walk",
	   Duration: 0,
	   StartTime: 0,
	   EndTime: 0,
	   },
	   },
	   })
	   return actionList
	   //ForPlural
	   action:=&Action{}
	   action.Set(&proto.Action{
	   ActionDescription: "",
	   Duration: 0,
	   StartTime: 0,
	   EndTime: 0,
	   })
	   return action
	*/
}
func (m *ActionManager) Current(d bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me, this is where you read this data from, could be connected to a database or a service
	panic("implement me")
}
func (m *ActionManager) SetServiceResponse(index uint64, response []byte, entity bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) {
	log := logger.GetLogger().WithField("ActionManager", "SetServiceResponse")
	protoActionList := &proto.ActionList{}
	err := proto1.Unmarshal(response, protoActionList)
	if err != nil {
		log.Errorf("Actions Props ByteStream Handled Error: %s", err)
	}
	actionList := &ActionList{}
	actionList.Set(protoActionList)
	ctx.SetResultData(actionList)
}
