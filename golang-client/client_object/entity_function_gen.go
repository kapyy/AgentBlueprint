package dobit_client

import (
	bpcontext "golang-client/bpcontext"
	logger "golang-client/modules/logger"
)

func (d *AgentEntity) InsertActionWithObservation(i *bpcontext.DataInstance, input bpcontext.UserInput) {
	queryctx := i.LinkNewQueryContext(i.Data())
	query := d.Queries[110000001]
	ctx := bpcontext.NewAgentContext(d, queryctx)
	queryctx.SetInputText(&input.InputText)
	query.BPFunctionNodes.FunctionParam.InputText = &input.InputText
	query.call(d, ctx)
}
func (d *AgentEntity) ActionFormatter(i *bpcontext.DataInstance) {
	log := logger.GetLogger().WithField("func", "ActionFormatter")
	byteData, err := i.Data().Marshal()
	if err != nil {
		log.Errorf("i.Data().Marshal error: %v", err)
		return
	}
	queryCtx := i.LinkNewQueryContext(i.Data())
	err = d.callSubordinateFunction(310000002, 1001, byteData, queryCtx)
	if err != nil {
		log.Errorf("callSubordinateFunction error: %v", err)
		return
	}
}
