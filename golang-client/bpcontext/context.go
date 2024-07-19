package bpcontext

import (
	"golang.org/x/net/context"
)

// TODO will be deprecated soon, all callback will be handeled together
type Callback struct {
	ResponseChannel ResponseChan[any]
	ResponseCancel  context.CancelFunc
}
type RootCache struct {
	DataProperties map[uint64]DataPropertyInterface
}

// -----------------------------------------------
// TODO need rework to middle step protoData extraction
type Memory struct {
	Labels []string
	Data   string
}

const (
	ContextInitialized   = iota //SetServiceResponse() Initialize
	ContextDataInProcess        //called a function_caller pipeline
	ContextDataActive           // Pipeline SetServiceResponse() Received pass on Context
	ContextDataComplete         // Mark Complete depends on children's Completion
	ContextDataInterrupted
	ContextDataBackLogged
)

// QueryContextInterface holds context protoData and response along a chain of pipelines
// A QueryContext node should be initialized and passed into an APM pipeline call
// If the pipeline is called at the end of another pipeline,
//
//	this initialized queryContext should be attached to that from previous pipeline
//
// All recorded protoData are stored in root node
type QueryContextInterface interface {
	//------Ctx Data Properties ------------
	InputText() *string
	SetInputText(inputText *string)
	InputData() DataPropertyInterface
	SetInputData(inputData DataPropertyInterface)
	RootCache() *RootCache
	SetRootCache(rootCache *RootCache)
	ResultData() []DataPropertyInterface
	SetResultData(resData DataPropertyInterface)
	SetListResultData(resData []DataPropertyInterface)
	TargetIndex() uint64
	SetTargetIndex(targetIndex uint64)
	//No Runtime Changes on the Structure
	Callback() *Callback
	//SetCallback(callback Callback)
	ParentQueryContext() QueryContextInterface
	//SetParentContext(parentContext QueryContextInterface)
	DataContextRoot() DataInstanceContextInterface
	//SetContextRoot(contextRoot DataInstanceContextInterface)
}

// DataInstanceContextInterface holds context tree of dataInstance and its derived protoData
// A DataInstanceContext node should be init when a new protoData instance is initialized with
// All derived dataInstances should attach their DataInstanceContext to root
type DataInstanceContextInterface interface {
	ForceDataContextInterrupt()
	CheckDataContextCompletion() bool
	MarkDataContextCompletion()
	UpdateDataContextCompletion()
	LinkNewQueryContext(input ...DataPropertyInterface) QueryContextInterface
	Callback() *Callback
	SetCallback(callback *Callback)

	//Data() *DataInstance
	//SetData(protoData *DataInstance)
	GetQueryContext() QueryContextInterface
	ClearQueryContext()
	TryLockContext() bool
	UnlockContext()
	NewBranchContext() DataInstanceContextInterface
	NewReferenceContext() DataInstanceContextInterface
	AddDataContextEventListener(eventId int, listener Listener)
}

//------------------------------------------------

type AgentContext struct {
	Dobit       AgentInterface
	DataContext QueryContextInterface
}

//	func FromContext(ctx *DataContext) *DataContext {
//		var dataContext *DataContext
//		if ctx == nil {
//			dataContext = &DataContext{
//				RootCache: &RootCache{
//					DataProperties: make(map[uint64]DataPropertyInterface),
//				},
//			}
//		} else {
//			dataContext = &DataContext{
//				ParentQueryContext: dataContext,
//			}
//		}
//		return dataContext
//	}
func NewAgentContext(d AgentInterface, ctx QueryContextInterface) *AgentContext {
	if ctx == nil {
		ctx = &FunctionContext{}
	}
	if ctx.RootCache() == nil {
		ctx.SetRootCache(&RootCache{
			DataProperties: make(map[uint64]DataPropertyInterface),
		})
	}

	return &AgentContext{
		Dobit:       d,
		DataContext: ctx,
	}
}

//func ContextWithCallback() QueryContextInterface {
//	return &FunctionContext{
//		inputText: nil,
//		inputData: nil,
//		rootCache: &RootCache{
//			DataProperties: make(map[uint64]DataPropertyInterface),
//		},
//		resultData:  nil,
//		targetIndex: 0,
//	}
//}

//	func NewDataContext() *DataContext {
//		return InitDataContext(&DataContext{})
//	}
//
//	func InitDataContext(ctx *DataContext) *DataContext {
//		ctx.DataProperties = make(map[uint64]DataPropertyInterface)
//		return ctx
//	}
