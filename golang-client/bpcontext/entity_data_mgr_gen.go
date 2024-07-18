package bpcontext

type ResponseChan[T interface{}] chan T

type DescriptorInterface interface {
	GetDescriptor(index uint64, d DobitInterface, ctx QueryContextInterface) DataPropertyInterface
	GetProps(list DataPropertyInterface, index uint64) ([]byte, string)
	//TODO Set SerivceResponse might able to be moved to Manager's Parser
	SetServiceResponse(index uint64, response []byte, entity DobitInterface, ctx QueryContextInterface)
}
type DataPropertyInterface interface {
	GetPropIndex(index uint64) (interface{}, string)
}

var MainServiceMgr = map[uint64]DescriptorInterface{}

func RegisterMgrComponent(index uint64, mgr DescriptorInterface) {
	MainServiceMgr[index] = mgr
}
func GetDataManager(index uint64) DescriptorInterface {
	return MainServiceMgr[index]
}

type DobitInterface interface {

	//Active PipelineFunctions

	InsertActionsWithObservation(datactx QueryContextInterface)

	GetDataManager(mgr int) DataManagerInterface
}
