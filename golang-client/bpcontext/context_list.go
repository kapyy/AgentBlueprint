package bpcontext

import "sync"

// root Context as in APM Function Node
type FunctionContext struct {
	inputText   *string
	inputData   DataPropertyInterface //temp cache, if all function input uses DataInstance Data, then this can be removed
	rootCache   *RootCache
	resultData  []DataPropertyInterface //temp cache, if all function output to DataInstance Data, then this can be removed
	targetIndex uint64

	contextRoot DataInstanceContextInterface
}

func (f *FunctionContext) InputText() *string {
	return f.inputText
}

func (f *FunctionContext) SetInputText(inputText *string) {
	f.inputText = inputText
}

func (f *FunctionContext) InputData() DataPropertyInterface {
	return f.inputData
}

func (f *FunctionContext) SetInputData(inputData DataPropertyInterface) {
	f.inputData = inputData
}

func (f *FunctionContext) RootCache() *RootCache {
	return f.rootCache
}

func (f *FunctionContext) SetRootCache(rootCache *RootCache) {
	f.rootCache = rootCache
}

func (f *FunctionContext) ResultData() []DataPropertyInterface {
	return f.resultData
}
func (f *FunctionContext) SetResultData(resData DataPropertyInterface) {
	f.resultData = []DataPropertyInterface{resData}
}

func (f *FunctionContext) SetListResultData(resData []DataPropertyInterface) {
	f.resultData = resData
}

func (f *FunctionContext) TargetIndex() uint64 {
	return f.targetIndex
}

func (f *FunctionContext) SetTargetIndex(targetIndex uint64) {
	f.targetIndex = targetIndex
}

func (f *FunctionContext) Callback() *Callback {
	return f.DataContextRoot().Callback()
}

//func (f *FunctionContext) SetCallback(callback Callback) {
//	f.DataContextRoot().SetCallback(callback)
//}

func (f *FunctionContext) ParentQueryContext() QueryContextInterface {
	return f
}

func (f *FunctionContext) DataContextRoot() DataInstanceContextInterface {
	return f.contextRoot
}

func newFunctionContext(root DataInstanceContextInterface) *FunctionContext {
	return &FunctionContext{
		rootCache: &RootCache{
			DataProperties: make(map[uint64]DataPropertyInterface),
		},
		contextRoot: root,
	}
}

// DataContext that does not pass out current pipeline
type DataNodeContext struct {
	//inputData     DataPropertyInterface
	//TODO warning, for now InputText are not used by protoData node over abstract protoData nodes, so overwriting is fine.
	//if needed, abstract string (inputData String format) need to be passed separately
	inputText     *string // InputText are overwritten by abstract Data
	parentContext QueryContextInterface
}

func (g *DataNodeContext) DataContextRoot() DataInstanceContextInterface {
	return g.parentContext.DataContextRoot()
}

func (g *DataNodeContext) ParentQueryContext() QueryContextInterface {
	return g.parentContext
}

func (g *DataNodeContext) InputText() *string {
	return g.inputText
}

func (g *DataNodeContext) SetInputText(inputText *string) {
	g.SetInputText(inputText)
}

func (g *DataNodeContext) InputData() DataPropertyInterface {
	return g.parentContext.InputData()
}

func (g *DataNodeContext) SetInputData(inputData DataPropertyInterface) {
	g.parentContext.SetInputData(inputData)
}

func (g *DataNodeContext) RootCache() *RootCache {
	return g.parentContext.RootCache()
}

func (g *DataNodeContext) SetRootCache(rootCache *RootCache) {
	g.parentContext.SetRootCache(rootCache)
}

func (g *DataNodeContext) ResultData() []DataPropertyInterface {
	return g.parentContext.ResultData()
}
func (g *DataNodeContext) SetResultData(resData DataPropertyInterface) {
	g.parentContext.SetResultData(resData)
}

func (g *DataNodeContext) SetListResultData(resData []DataPropertyInterface) {
	g.parentContext.SetListResultData(resData)
}

func (g *DataNodeContext) TargetIndex() uint64 {
	return g.parentContext.TargetIndex()
}

func (g *DataNodeContext) SetTargetIndex(targetIndex uint64) {
	g.parentContext.SetTargetIndex(targetIndex)
}

func (g *DataNodeContext) Callback() *Callback {
	return g.parentContext.Callback()
}

//	func (g *DataNodeContext) SetCallback(callback Callback) {
//		g.parentContext.SetCallback(callback)
//	}
func NewDataContext(parent QueryContextInterface) *DataNodeContext {
	return &DataNodeContext{
		parentContext: parent,
	}
}

// Function Context that pass through pipeline
// not sure if we need offgridContext to pass on
type StructuralContext struct {
	ContextEventManager
	proceduralTag int
	queryContext  QueryContextInterface

	//ParentStructure: Completion is passed upwards, for branching style like event-> actions
	//when all children completed, parent will be marked as completed
	childrenContext []DataInstanceContextInterface
	parentContext   DataInstanceContextInterface

	//ReferenceStructure: Completion is passed downwards with event, for reference style like observation -> actions
	//when parent completed, all children will be marked as completed
	//TODO Context force DataInstance to interrupt must be implemented here
	referenceContext DataInstanceContextInterface

	callback *Callback //TODO might be able to be merged with DataInstance.DistributedEvent
	mu       sync.RWMutex
}

//func (s *StructuralContext) Data() *DataInstance {
//	return s.protoData
//}
//
//func (s *StructuralContext) SetData(protoData *DataInstance) {
//	s.protoData = protoData
//}

func (s *StructuralContext) TryLockContext() bool {
	return s.mu.TryLock()
}
func (s *StructuralContext) UnlockContext() {
	s.mu.Unlock()
}

func (s *StructuralContext) Callback() *Callback {
	return s.callback
}

func (s *StructuralContext) SetCallback(callback *Callback) {
	s.callback = callback
}

// To interrupt as Completion
func (s *StructuralContext) ForceDataContextInterrupt() {
	s.proceduralTag = ContextDataInterrupted
	s.Invoke(ContextDataInterrupted)
	for _, child := range s.childrenContext {
		child.ForceDataContextInterrupt()
	}
}

func (s *StructuralContext) MarkDataContextCompletion() {
	s.proceduralTag = ContextDataComplete
	s.Invoke(ContextDataComplete)
	s.parentContext.UpdateDataContextCompletion()
}

func (s *StructuralContext) UpdateDataContextCompletion() {
	for _, child := range s.childrenContext {
		if !child.CheckDataContextCompletion() {
			return
		}
	}
	s.MarkDataContextCompletion()
}
func (s *StructuralContext) CheckDataContextCompletion() bool {
	//Interrupted  are considered as completed?
	return s.proceduralTag >= ContextDataComplete
}
func (s *StructuralContext) LinkNewQueryContext(input ...DataPropertyInterface) QueryContextInterface {
	funcCtx := newFunctionContext(s)
	if len(input) > 0 {
		funcCtx.SetInputData(input[0])
	}
	s.queryContext = funcCtx
	return funcCtx
}
func (s *StructuralContext) GetQueryContext() QueryContextInterface {
	return s.queryContext
}
func (s *StructuralContext) ClearQueryContext() {
	s.queryContext = nil
}

func (s *StructuralContext) NewBranchContext() DataInstanceContextInterface {
	branchCtx := &StructuralContext{
		callback:      s.callback,
		proceduralTag: ContextInitialized,
		parentContext: s,
	}
	s.childrenContext = append(s.childrenContext, branchCtx)
	return branchCtx
}
func (s *StructuralContext) NewReferenceContext() DataInstanceContextInterface {
	branchCtx := &StructuralContext{
		callback:         s.callback,
		proceduralTag:    ContextInitialized,
		referenceContext: s,
	}
	s.AddListener(ContextDataComplete, func(args ...any) {
		//Does Parent Context force children to interrupt?
		branchCtx.ForceDataContextInterrupt()
	})
	return branchCtx
}
func newContext() *StructuralContext {
	return &StructuralContext{
		proceduralTag: ContextInitialized,
	}
}

func (s *StructuralContext) AddDataContextEventListener(eventId int, listener Listener) {
	s.AddListener(eventId, listener)
}
