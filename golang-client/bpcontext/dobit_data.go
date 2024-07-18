package bpcontext

import "sync"

type DataInstance struct {
	*StructuralData
	DataInstanceContextInterface
}

type StructuralData struct {
	//StructuralData currently equals to DataContext.FunctionContext.ResultData
	//TODO might able to merge them together (Might able to merge after SetServiceResponse are moved to processor)
	dataMu       *sync.RWMutex
	DataCond     *sync.Cond
	data         DataPropertyInterface
	dataStateTag int

	DistributedEvent
}

func (d *StructuralData) Data() DataPropertyInterface {
	return d.data
}

func (d *StructuralData) SetData(structuralData DataPropertyInterface) {
	d.data = structuralData
}

func (d *StructuralData) DataStateTag() int {
	return d.dataStateTag
}

func (d *StructuralData) SetDataStateTag(dataStateTag int) {
	d.dataStateTag = dataStateTag
	d.Invoke(dataStateTag)
}

func newStructuralData(data DataPropertyInterface) *StructuralData {
	lock := &sync.RWMutex{}
	return &StructuralData{
		dataMu:       lock,
		DataCond:     sync.NewCond(lock),
		data:         data,
		dataStateTag: 0,
	}
}
func NewDataInstance(data DataPropertyInterface) *DataInstance {
	return &DataInstance{
		StructuralData:               newStructuralData(data),
		DataInstanceContextInterface: newContext(),
	}
}
func NewDataInstanceWithContext(data DataPropertyInterface, context DataInstanceContextInterface) *DataInstance {
	return &DataInstance{
		StructuralData:               newStructuralData(data),
		DataInstanceContextInterface: context,
	}
}

type DataProcessorInterface interface {
	Process(entity DobitInterface, instances ...*DataInstance)
}

type DataManagerInterface interface {
	Init(dobit DobitInterface)
	InsertDataInstance(instance *DataInstance)
	Update(DobitInterface)
	GetProcessor(tag int) func(entity DobitInterface, ctx ...*DataInstance)
	Event() InternalEventInterface
	AddListener(eventType int, listener Listener)
	GetFirstTaggedDataInstance(tag ...int) *DataInstance
	GetTaggedDataInstances(tag ...int) []*DataInstance
}
