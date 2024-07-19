package bpcontext

type AgentInterface interface {
	InsertActionWithObservation(i *DataInstance, input UserInput)
	ActionFormatter(i *DataInstance)
	GetDataManager(mgr int) DataManagerInterface
}
