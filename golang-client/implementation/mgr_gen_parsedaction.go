package implementation

import (
	bpcontext "golang-client/bpcontext"
	protodata "golang-client/message/protoData"
	logger "golang-client/modules/logger"
	proto "google.golang.org/protobuf/proto"
)

type ParsedActionManager struct{}

func (m *ParsedActionManager) GetProps(protoData bpcontext.DataPropertyInterface, index uint64) ([]byte, string) {
	log := logger.GetLogger().WithField("func", "ParsedActionManagerGetProps")
	interfaceObj, stringObj := protoData.GetPropIndex(index)
	protoObj, ok := interfaceObj.(*protodata.ParsedAction)
	if !ok {
		log.Debugf("Conversion failed. The interface does not hold a * %v.", "ParsedAction")
	}
	byteStream, err := proto.Marshal(protoObj)
	if err != nil {
		log.Errorf("ParsedActions Props ByteStream Handled Error: %v", err)
	}
	return byteStream, stringObj
}
func (m *ParsedActionManager) GetDescriptor(index uint64, d bpcontext.DobitInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	log := logger.GetLogger().WithField("func", "ParsedActionManagerGetDescriptor")
	switch index {
	case 0:
		return m.Default(d, ctx)
	default:
		log.Errorf("Descriptor not found: %d", index)
		return m.Default(d, ctx)
	}
}
