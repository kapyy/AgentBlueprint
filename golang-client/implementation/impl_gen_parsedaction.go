package implementation

import bpcontext "golang-client/bpcontext"

func (m *ParsedActionManager) Default(d bpcontext.DobitInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me
	panic("implement me")
	/*
	   ml := MemoryLog{startTime: 100, endTime: 200, actionDescription: "test"}

	   	return DataPropertyInterface{&ml}
	*/
}
func (m *ParsedActionManager) SetServiceResponse(index uint64, response []byte, entity bpcontext.DobitInterface, ctx bpcontext.QueryContextInterface) {
	// Currently Not In Use
	/*
		switch index {
		case 3001: //MemoryDistillToLongTerm
			protoData := protoData.MemoryLogs{}
			err := proto.Unmarshal(response, &protoData)
			if err != nil {
				log.Fatal("MemoryLogs Props ByteStream Handled Error: ", err)
			}
			fmt.Print("MemoryDistillToLongTerm..MemoryLogs:", protoData)
		case 4001: //SummarizeAgent
			protoData := protoData.MemoryLogs{}
			err := proto.Unmarshal(response, &protoData)
			if err != nil {
				log.Fatal("MemoryLogs Props ByteStream Handled Error: ", err)
			}
			fmt.Print("SummarizeAgent..MemoryLogs:", protoData)

		}
	*/
}
