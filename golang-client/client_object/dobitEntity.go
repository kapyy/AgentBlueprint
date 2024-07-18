package dobit_client

import (
	"golang-client/bpcontext"
	"golang-client/implementation"
	"golang-client/modules/logger"
	"google.golang.org/protobuf/proto"
)

type DobitEntity struct {
	dataManager map[int]bpcontext.DataManagerInterface
	Queries     map[uint64]Query
}

func (d *DobitEntity) GetDataManager(mgr int) bpcontext.DataManagerInterface {
	return d.dataManager[mgr]
}

func (d *DobitEntity) InsertActionsWithObservation(datactx bpcontext.QueryContextInterface) {
	query := d.Queries[310400400]
	ctx := bpcontext.Context(d, datactx)
	query.call(d, ctx)
}

func (d *DobitEntity) ActionFormatter(datactx bpcontext.QueryContextInterface, action bpcontext.DataPropertyInterface) {
	action_log, ok := action.(*implementation.Action)
	if !ok {
		logger.GetLogger().Errorf("ActionFormatter: protoData type error")
	}
	byteData, _ := proto.Marshal(action_log.Action)

	datactx.SetInputData(action_log)
	datactx.SetTargetIndex(0)
	err := d.callSubordinateFunction(400100300, 1007, byteData, datactx)
	if err != nil {
		logger.GetLogger().WithField("func", "ActionFormatter").Errorf("callSubordinateFunction error: %v", err)
	}
}
