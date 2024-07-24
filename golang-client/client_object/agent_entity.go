package dobit_client

import (
	"golang-client/bpcontext"
	message "golang-client/message/proto"
	"golang-client/modules/logger"
	"google.golang.org/protobuf/proto"
	"os"
)

type AgentEntity struct {
	dataManager map[int]bpcontext.DataManagerInterface
	Queries     map[uint64]Query
}

func (d *AgentEntity) GetDataManager(mgr int) bpcontext.DataManagerInterface {
	return d.dataManager[mgr]
}
func DeserializeAPMToEntity(filename string, entity *AgentEntity) {
	log := logger.GetLogger().WithField("func", "DeserializeAPMToEntity")
	in, err := os.ReadFile(filename)
	if err != nil {
		log.Errorf("ReadFile error %v", err.Error())
	}
	object := message.ApmFile{}
	err = proto.Unmarshal(in, &object)
	if err != nil {
		log.Errorf("Unmarshal error %v", err.Error())
	}

	for _, tree := range object.Trees {
		//if tree.TreeType == 0 {
		//	log.Fatal("DataNode cannot be root node")
		//}
		entity.registerFunctionCall(tree.RootNode)
	}

}
