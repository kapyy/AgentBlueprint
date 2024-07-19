package dobit_client

import (
	"golang-client/bpcontext"
)

type AgentEntity struct {
	dataManager map[int]bpcontext.DataManagerInterface
	Queries     map[uint64]Query
}

func (d *AgentEntity) GetDataManager(mgr int) bpcontext.DataManagerInterface {
	return d.dataManager[mgr]
}
