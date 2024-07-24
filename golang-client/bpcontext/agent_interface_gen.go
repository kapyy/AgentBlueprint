package bpcontext

import implementation "golang-client/implementation"

type AgentInterface interface {
	InsertActionWithObservation(i *DataInstance, input UserInput) *implementation.ActionList
	ActionFormatter(i *DataInstance) *implementation.ParsedAction
	GetDataManager(mgr int) DataManagerInterface
}
