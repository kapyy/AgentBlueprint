package implementation

import (
	bpcontext "golang-client/bpcontext"
	proto "golang-client/message/proto"
	logger "golang-client/modules/logger"
	proto1 "google.golang.org/protobuf/proto"
)

func (m *ParsedActionManager) Default(d bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me, this is where you read this data from, could be connected to a database or a service
	panic("implement me")
	/*
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
func (m *ParsedActionManager) SetServiceResponse(index uint64, response []byte, entity bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) {
	log := logger.GetLogger().WithField("ParsedActionManager", "SetServiceResponse")
	protoParsedAction := &proto.ParsedAction{}
	err := proto1.Unmarshal(response, protoParsedAction)
	if err != nil {
		log.Errorf("ParsedAction Props ByteStream Handled Error: %s", err)
	}
	parsedaction := &ParsedAction{}
	parsedaction.Set(protoParsedAction)
	ctx.SetResultData(parsedaction)
}
