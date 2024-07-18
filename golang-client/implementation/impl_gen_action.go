package implementation

import bpcontext "golang-client/bpcontext"

func (m *ActionManager) Default(d bpcontext.DobitInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me
	panic("implement me")
	/*
	   bs := BaseObject{fulldescription: ""}

	   	return &BaseObjects{baseobjects: []*BaseObject{&bs}}
	*/
}
func (m *ActionManager) Current(d bpcontext.DobitInterface, ctx bpcontext.QueryContextInterface) bpcontext.DataPropertyInterface {
	// TODO implement me
	panic("implement me")
}
func (m *ActionManager) SetServiceResponse(index uint64, response []byte, entity bpcontext.DobitInterface, ctx bpcontext.QueryContextInterface) {
	// TODO implement me
	panic("implement me")
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
