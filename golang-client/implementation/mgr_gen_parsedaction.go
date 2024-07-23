package implementation

import (
	bpcontext "golang-client/bpcontext"
	proto "golang-client/message/proto"
	logger "golang-client/modules/logger"
	proto1 "google.golang.org/protobuf/proto"
)

type ParsedActionManager struct{}

func (m *ParsedActionManager) GetProps(proto bpcontext.DataPropertyInterface, index uint64) ([]byte, string) {
	log := logger.GetLogger().WithField("func", "ParsedActionManagerGetProps")
	interfaceObj, stringObj := proto.GetPropIndex(index)
	protoObj, ok := interfaceObj.(*proto.ParsedAction)
	if !ok {
		log.Debugf("Conversion failed. The interface does not hold a * %v.", "ParsedAction")
	}
	byteStream, err := proto1.Marshal(protoObj)
	if err != nil {
		log.Errorf("ParsedActions Props ByteStream Handled Error: %v", err)
	}
	return byteStream, stringObj
}
func (m *ParsedActionManager) GetDescriptor(index uint64, d bpcontext.AgentInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	log := logger.GetLogger().WithField("func", "ParsedActionManagerGetDescriptor")
	switch index {
	case 0:
		return m.Default(d, ctx)
	default:
		log.Errorf("Descriptor not found: %d", index)
		return m.Default(d, ctx)
	}
}
