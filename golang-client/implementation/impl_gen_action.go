package implementation

import (
	bpcontext "golang-client/bpcontext"
	protodata "golang-client/message/protoData"
	logger "golang-client/modules/logger"
	proto "google.golang.org/protobuf/proto"
)

func (m *ActionManager) Default(d bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me, this is where you read this data from, could be connected to a database or a service
	panic("implement me")
	/*
	   actions := &ActionList{}
	   actionList.Set(&protodata.ActionList{
	   Actions:[]*protodata.Action{
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
	   action.Set(&protodata.Action{
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
	protoActionList := &protodata.ActionList{}
	err := proto.Unmarshal(response, protoActionList)
	if err != nil {
		log.Errorf("Actions Props ByteStream Handled Error: %s", err)
	}
	actionList := &ActionList{}
	actionList.Set(protoActionList)
	ctx.SetResultData(actionList)
}
