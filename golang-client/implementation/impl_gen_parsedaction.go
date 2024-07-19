package implementation

import (
	bpcontext "golang-client/bpcontext"
	protodata "golang-client/message/protoData"
	logger "golang-client/modules/logger"
	proto "google.golang.org/protobuf/proto"
)

func (m *ParsedActionManager) Default(d bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me, this is where you read this data from, could be connected to a database or a service
	panic("implement me")
	/*
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
func (m *ParsedActionManager) SetServiceResponse(index uint64, response []byte, entity bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) {
	log := logger.GetLogger().WithField("ParsedActionManager", "SetServiceResponse")
	protoParsedAction := &protodata.ParsedAction{}
	err := proto.Unmarshal(response, protoParsedAction)
	if err != nil {
		log.Errorf("ParsedAction Props ByteStream Handled Error: %s", err)
	}
	parsedaction := &ParsedAction{}
	parsedaction.Set(protoParsedAction)
	ctx.SetResultData(parsedaction)
}
