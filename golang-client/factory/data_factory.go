package factory

import (
	"golang-client/bpcontext"
	"golang-client/modules/logger"
)

func deserializeData(dataID uint64, ctx *bpcontext.AgentContext) ([]byte, string) {
	log := logger.GetLogger().WithField("func", "deserializeData").WithField("dataID", dataID)
	dataType, descriptor, property := DisassembleDataID(dataID)
	dataMgr := bpcontext.GetDataManager(dataType)
	if dataMgr == nil {
		log.Errorf("DataMgr for dataId %d is nil", dataID)
	}
	cacheIndex := dataType*100 + descriptor

	var data bpcontext.DataPropertyInterface
	//Abstract dataType does not goes in cache
	if dataType == 4015 || dataType == 4022 || dataType == 4023 || dataType == 4027 {
		data = dataMgr.GetDescriptor(descriptor, ctx.Dobit, ctx.DataContext)
	} else {
		//cache_data, exist := properties_map[cache_index]
		if cacheData, exist := ctx.DataContext.RootCache().DataProperties[cacheIndex]; !exist {
			data = dataMgr.GetDescriptor(descriptor, ctx.Dobit, ctx.DataContext)
			ctx.DataContext.RootCache().DataProperties[cacheIndex] = data
		} else {
			data = cacheData
		}
	}
	rawData, propString := dataMgr.GetProps(data, property)

	//node := &apm.NodeData{
	//	//Data Node has no further use, so will not be implemented
	//}
	shortenStr := propString
	if len(shortenStr) > 20 {
		shortenStr = shortenStr[:20]
	}
	log.Debugf("DataID: %d, Property: %s ...", dataID, shortenStr)
	return rawData, propString
}

func DisassembleDataID(id uint64) (uint64, uint64, uint64) {
	return (id / 10000) % 10000, (id / 100) % 100, id % 100
}
