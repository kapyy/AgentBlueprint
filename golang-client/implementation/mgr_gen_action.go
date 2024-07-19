package implementation

import (
	bpcontext "golang-client/bpcontext"
	protodata "golang-client/message/protoData"
	logger "golang-client/modules/logger"
	proto "google.golang.org/protobuf/proto"
)

type ActionManager struct {
	ActionList ActionList
}

func (m *ActionManager) GetDescriptor(index uint64, d bpcontext.DobitInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	log := logger.GetLogger().WithField("func", "GetDescriptor")
	switch index {
	case 0:
		return m.Default(d, ctx)
	case 20:
		return m.Current(d, ctx)
	default:
		log.Errorf("No such Descriptor in Action Mgr")
		return m.Default(d, ctx)
	}
}
func (m *ActionManager) GetProps(list bpcontext.DataPropertyInterface, index uint64) ([]byte, string) {
	log := logger.GetLogger().WithField("ActionManager", "GetProps")
	listStruct, ok := list.(*ActionList)
	if !ok {
		log.Debugf("Conversion failed.GetProps List does not hold a *Action")
	}
	interfaceObj, stringObj := listStruct.GetPropIndex(index)
	serializeObj, ok := interfaceObj.(*protodata.ActionList)
	if !ok {
		log.Debugf("Conversion failed.GetProps Return does not hold a *Action")
	}
	byteStream, err := proto.Marshal(serializeObj)
	if err != nil {
		log.Errorf("Actions Props ByteStream Handled Error: %v", err)
	}
	return byteStream, stringObj
}
