package tools

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

func extDataGen(conf *DataYamlConfig, overwrite bool) {
	sortedExternalData := SortData(conf.ExternalData)
	for i, externalData := range sortedExternalData {
		name := externalData.Key
		print("Generating " + name + "...\n")
		if overwrite {
			f_descp := jen.NewFilePathName("implementation", "implementation")
			sysDataDescGen(f_descp, i, externalData, false)
			err := f_descp.Save("./implementation/desc_gen_" + strings.ToLower(name) + ".go")
			if err != nil {
				return
			}
			f_impl := jen.NewFilePathName("implementation", "implementation")
			extDataImplGen(f_impl, name, conf)
			err = f_impl.Save("./implementation/impl_gen_" + strings.ToLower(name) + ".go")
			if err != nil {
				return
			}
		}
		f_mgr := jen.NewFilePathName("implementation", "implementation")
		extDataMgrGen(f_mgr, name, conf)
		err := f_mgr.Save("./implementation/mgr_gen_" + strings.ToLower(name) + ".go")
		if err != nil {
			return
		}
	}
	//for name := range conf.ExternalData {
	//	print("Generating " + name + "...\n")
	//	if overwrite {
	//		f_descp := jen.NewFilePathName("implementation", "implementation")
	//		sysDataDescGen(f_descp, name, conf.ExternalData, false)
	//		err := f_descp.Save("./implementation/desc_gen_" + strings.ToLower(name) + ".go")
	//		if err != nil {
	//			return
	//		}
	//		f_impl := jen.NewFilePathName("implementation", "implementation")
	//		extDataImplGen(f_impl, name, conf)
	//		err = f_impl.Save("./implementation/impl_gen_" + strings.ToLower(name) + ".go")
	//		if err != nil {
	//			return
	//		}
	//	}
	//	f_mgr := jen.NewFilePathName("implementation", "implementation")
	//	extDataMgrGen(f_mgr, name, conf)
	//	err := f_mgr.Save("./implementation/mgr_gen_" + strings.ToLower(name) + ".go")
	//	if err != nil {
	//		return
	//	}
	//}
}

func extDataImplGen(impl *jen.File, name string, conf *DataYamlConfig) {
	impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("Default").Params(jen.Id("d").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
		jen.Comment("TODO implement me, this is where you read this data from, could be connected to a database or a service"),
		jen.Panic(jen.Lit("implement me")),

		jen.Comment("action:=&Action{}\n"+"action.Set(&proto.Action{\n"+"ActionDescription: \"\",\n"+"Duration: 0,\n"+"StartTime: 0,\n"+"EndTime: 0,\n"+"})\n"+"return action"),
	)
	for _, desc := range conf.ExternalData[name].Desc {
		if conf.Descriptor[desc] != nil {
			for desc_name := range conf.Descriptor[desc] {
				impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id(desc_name).Params(jen.Id("d").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
					jen.Comment("TODO implement me, this is where you read this data from, could be connected to a database or a service"),
					jen.Panic(jen.Lit("implement me")),
				)
			}
		}
	}
	impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("SetServiceResponse").Params(jen.Id("index").Uint64(), jen.Id("response").Id("[]byte"), jen.Id("entity").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).BlockFunc(func(g *jen.Group) {
		g.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit(name+"Manager"), jen.Lit("SetServiceResponse"))
		g.Id("proto"+name).Op(":=").Op("&").Qual("golang-client/message/proto", name).Values()
		g.Id("err").Op(":=").Qual("google.golang.org/protobuf/proto", "Unmarshal").Call(jen.Id("response"), jen.Id("proto"+name))
		g.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("log").Dot("Errorf").Call(jen.Lit(name+" Props ByteStream Handled Error: %s"), jen.Id("err")),
		)
		g.Id(strings.ToLower(name)).Op(":=").Op("&").Id(name).Values()
		g.Id(strings.ToLower(name)).Dot("Set").Call(jen.Id("proto" + name))
		g.Id("ctx").Dot("SetResultData").Call(jen.Id(strings.ToLower(name)))
	})
	// fmt.Printf("extDataImplGen: %v\n", impl)
}
func extDataMgrGen(mgr *jen.File, name string, conf *DataYamlConfig) {
	mgr.Type().Id(name + "Manager").Struct(
	//jen.Id(name + "s").Id("[]" + name),
	)
	mgr.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("GetProps").Params(jen.Id("proto").
		Qual("golang-client/bpcontext", "DataPropertyInterface"), jen.Id("index").Uint64()).Params(jen.Id("[]byte"), jen.Id("string")).Block(
		jen.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit("func"), jen.Lit(name+"ManagerGetProps")),
		jen.Id("interfaceObj, stringObj").Op(":=").Id("proto").Dot("GetPropIndex").Call(jen.Id("index")),
		jen.Id("protoObj, ok").Op(":=").Id("interfaceObj").Assert(jen.Op("*").Qual("golang-client/message/proto", name)),
		jen.If(jen.Id("!ok")).Block(
			jen.Id("log").Dot("Debugf").Call(jen.Lit("Conversion failed. The interface does not hold a * %v."), jen.Lit(name))),
		jen.Id("byteStream, err").Op(":=").Qual("google.golang.org/protobuf/proto", "Marshal").Call(jen.Id("protoObj")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("log").Dot("Errorf").Call(jen.Lit(name+"s Props ByteStream Handled Error: %v"), jen.Id("err"))),
		jen.Return(jen.Id("byteStream"), jen.Id("stringObj")),
	)
	mgr.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("GetDescriptor").Params(jen.Id("index").Uint64(), jen.Id("d").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
		jen.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit("func"), jen.Lit(name+"ManagerGetDescriptor")),
		jen.Switch(jen.Id("index")).BlockFunc(func(g *jen.Group) {
			g.Case(jen.Lit(0)).Block(
				jen.Return(jen.Id("m").Dot("Default").Params(jen.Id("d"), jen.Id("ctx"))),
			)
			for _, desc := range conf.ExternalData[name].Desc {
				if conf.Descriptor[desc] != nil {
					sortedDataDescriptorProperties := SortDescriptorProperty(conf.Descriptor[desc])
					for _, property := range sortedDataDescriptorProperties {
						desc_name := property.Key
						desc_i := property.Value
						g.Case(jen.Lit(desc_i)).Block(
							jen.Return(jen.Id("m").Dot(desc_name).Params(jen.Id("d"), jen.Id("ctx"))),
						)
					}
				}
			}
			g.Default().Block(
				jen.Id("log").Dot("Errorf").Call(jen.Lit("Descriptor not found: %d"), jen.Id("index")),
				jen.Return(jen.Id("m").Dot("Default").Params(jen.Id("d"), jen.Id("ctx"))),
			)
		}),
	)
}
