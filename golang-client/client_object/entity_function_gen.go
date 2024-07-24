package dobit_client

import (
	bpcontext "golang-client/bpcontext"
	implementation "golang-client/implementation"
	logger "golang-client/modules/logger"
)

func (d *AgentEntity) InsertActionWithObservation(i *bpcontext.DataInstance, input bpcontext.UserInput) *implementation.ActionList {
	log := logger.GetLogger().WithField("func", "InsertActionWithObservation")
	queryctx := i.LinkNewQueryContext(i.Data())
	query := d.Queries[110000001]
	ctx := bpcontext.NewAgentContext(d, queryctx)
	queryctx.SetInputText(&input.InputText)
	query.BPFunctionNodes.FunctionParam.InputText = &input.InputText
	query.call(d, ctx)
	actionlist, ok := queryctx.ResultData().(*implementation.ActionList)
	if !ok {
		log.Errorf("InsertActionWithObservation: Return data type error")
	}
	return actionlist
}
func (d *AgentEntity) ActionFormatter(i *bpcontext.DataInstance) *implementation.ParsedAction {
	log := logger.GetLogger().WithField("func", "ActionFormatter")
	byteData, err := i.Data().Marshal()
	if err != nil {
		log.Errorf("i.Data().Marshal error: %v", err)
		return nil
	}
	queryctx := i.LinkNewQueryContext(i.Data())
	err = d.callSubordinateFunction(310000002, 1001, byteData, queryctx)
	if err != nil {
		log.Errorf("callSubordinateFunction error: %v", err)
		return nil
	}
	parsedaction, ok := queryctx.ResultData().(*implementation.ParsedAction)
	if !ok {
		log.Errorf("ActionFormatter: Return data type error")
	}
	return parsedaction
}
