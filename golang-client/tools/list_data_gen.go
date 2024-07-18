package tools

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

func sysDataGen(conf *YamlConfig, overwrite bool) {
	sortedSystemData := SortData(conf.SystemData)
	for i, data := range sortedSystemData {
		name := data.Key
		print("Generating " + name + "...\n")
		if overwrite {
			f_descp := jen.NewFilePathName("implementation", "implementation")
			sysDataDescGen(f_descp, i, data, true)
			err := f_descp.Save("./implementation/desc_gen_" + strings.ToLower(name) + ".go")
			if err != nil {
				return
			}
			f_impl := jen.NewFilePathName("implementation", "implementation")
			sysDataImplGen(f_impl, name, conf)
			err = f_impl.Save("./implementation/impl_gen_" + strings.ToLower(name) + ".go")
			if err != nil {
				return
			}
		}
		f_mgr := jen.NewFilePathName("implementation", "implementation")
		sysDataMgrGen(f_mgr, name, conf)
	}

}

//func sysDataGen(conf *YamlConfig, overwrite bool) {
//	//sortedSystemData := SortData(conf.SystemData)
//	for name := range conf.SystemData {
//		print("Generating " + name + "...\n")
//		if overwrite {
//			f_descp := jen.NewFilePathName("implementation", "implementation")
//			sysDataDescGen(f_descp, name,conf.SystemData , true)
//			err := f_descp.Save("./implementation/desc_gen_" + strings.ToLower(name) + ".go")
//			if err != nil {
//				return
//			}
//			f_impl := jen.NewFilePathName("implementation", "implementation")
//			sysDataImplGen(f_impl, name, conf)
//			err = f_impl.Save("./implementation/impl_gen_" + strings.ToLower(name) + ".go")
//			if err != nil {
//				return
//			}
//		}
//		f_mgr := jen.NewFilePathName("implementation", "implementation")
//		sysDataMgrGen(f_mgr, name, conf)
//	}
//
//}

func sysDataDescGen(descp *jen.File, i int, dataEntry DataEntry, hasPlural bool) {
	name := dataEntry.Key
	sortedProperties := SortDataProerty(dataEntry.Value)
	descp.Type().Id(dataEntry.Key).StructFunc(func(g *jen.Group) {
		for _, property := range sortedProperties {
			if isEmbeddedProperty(property.Value.Type) {
				g.Id("embedded_" + strings.ToLower(property.Key)).Id(convertIfReference(property.Value.Type))
			}

		}
		g.Op("*").Qual("golang-client/message/protoData", name)
	})
	descp.Func().Params(jen.Id("s").Id("*"+dataEntry.Key)).Id("Default").Call().Id("*").Qual("golang-client/message/protoData", name).Block(
		//jen.Return(jen.Op("&").Qual("golang-client/message/protoData", name).Values(jen.DictFunc(func(d jen.Dict) {
		//	for _, property := range sortedProperties {
		//		d[jen.Id(property.Key)] = jen.Id("s").Dot(property.Key).Params()
		//	}
		//}))),
		jen.Return(jen.Id("s").Dot(name)),
	)
	descp.Func().Params(jen.Id("s").Id("*" + dataEntry.Key)).Id("Set").Params(jen.Id(strings.ToLower(name)).Id("*").Qual("golang-client/message/protoData", name)).Block(
		jen.Id("s").Dot(name).Op("=").Id(strings.ToLower(name)),
	)
	descp.Func().Params(jen.Id("s").Id("*" + name)).Id("FullString").Call().Id("string").BlockFunc(func(g *jen.Group) {
		var return_val = jen.Return()
		firstprop := true
		for _, property := range sortedProperties {
			if !firstprop {
				return_val.Id("+")
			}
			return_val.Id("s").Dot(property.Key + "String").Call()
			firstprop = false
		}
		g.Add(return_val)
	})
	for _, property := range sortedProperties {
		prop_name := property.Key
		prop := property.Value
		descp.Func().Params(jen.Id("s").Id("*" + name)).Id(prop_name).Call().Add(convertProtoPropPtr(prop.Type)).Block(
			//jen.Comment("TODO: implement me"),
			//jen.Panic(jen.Lit("implement me")),
			//jen.Return(jen.Add(convertProtoPropObj(prop.Type))),
			jen.ReturnFunc(func(g *jen.Group) {
				if isEmbeddedProperty(property.Value.Type) {
					g.Id("s").Dot("embedded_" + strings.ToLower(property.Key))
				} else {
					g.Id("s").Dot(name).Dot(prop_name)
				}

			}),
		)
		descp.Func().Params(jen.Id("s").Id("*" + name)).Id("Set" + prop_name).Params(jen.Id(strings.ToLower(prop_name)).Id(convertIfReferenceWithGrpc(prop.Type))).BlockFunc(func(g *jen.Group) {
			if isEmbeddedProperty(property.Value.Type) {
				g.Id("s").Dot("embedded_" + strings.ToLower(property.Key)).Dot("Set").Params(jen.Id(strings.ToLower(prop_name)))
			}
			g.Id("s").Dot(name).Dot(prop_name).Op("=").Id(strings.ToLower(prop_name))
		})
		descp.Func().Params(jen.Id("s").Id("*"+name)).Id(prop_name+"String").Call().Id("string").Block(
			jen.Comment("TODO: implement me"),
			jen.Panic(jen.Lit("implement me")),
		)
	}
	descp.Func().Params(jen.Id("s").Id("*"+name)).Id("GetPropIndex").Params(jen.Id("index").Uint64()).Params(jen.Id("interface{}"), jen.Id("string")).Block(
		jen.Switch(jen.Id("index")).BlockFunc(func(g *jen.Group) {
			g.Case(jen.Id("0")).Block(
				jen.Return(jen.Id("s").Dot("Default").Call(), jen.Id("s").Dot("FullString").Call()),
			)
			for _, property := range sortedProperties {
				prop_name := property.Key
				prop := property.Value
				g.Case(jen.Id(fmt.Sprintf("%d", prop.Index))).Block(
					jen.ReturnFunc(func(g *jen.Group) {
						if isEmbeddedProperty(property.Value.Type) {
							g.Id("&").Qual("golang-client/message/protoData", name).Values(jen.Dict{
								jen.Id(prop_name): jen.Id("s").Dot(prop_name).Call().Dot(prop.Type),
							}).Id(",").Id("s").Dot(prop_name + "String").Call()
						} else {
							g.Id("&").Qual("golang-client/message/protoData", name).Values(jen.Dict{
								jen.Id(prop_name): jen.Id("s").Dot(prop_name).Call(),
							}).Id(",").Id("s").Dot(prop_name + "String").Call()
						}

					}),
				)
			}
			g.Default().Block(
				jen.Return(jen.Id("&").Qual("golang-client/message/protoData", name).Values(), jen.Lit("")),
			)
		}),
	)

	if hasPlural {
		descp.Type().Id(name + "s").Struct(
			jen.Id(strings.ToLower(name + "s")).Id("[]*").Id(name),
		)
		descp.Func().Params(jen.Id("sl").Id("*"+name+"s")).Id("GetPropIndex").Params(jen.Id("index").Uint64()).Params(jen.Id("interface{}"), jen.Id("string")).Block(
			jen.Id("protoList").Op(":=").Make(jen.Index().Id("*").Qual("golang-client/message/protoData", name), jen.Len(jen.Id("sl").Dot(strings.ToLower(name+"s")))),
			jen.Id("stringList").Op(":=").Lit(""),
			jen.For(jen.Id("i,").Id("s").Op(":=").Range().Id("sl").Dot(strings.ToLower(name+"s"))).Block(
				jen.Id("protoObj").Op(",").Id("stringObj").Op(":=").Id("s").Dot("GetPropIndex").Call(jen.Id("index")),
				jen.Id("protoList").Index(jen.Id("i")).Op("=").Id("protoObj").Assert(jen.Op("*").Qual("golang-client/message/protoData", name)),
				jen.Id("stringList").Op("+=").Id("stringObj"),
			),
			jen.Return(jen.Id("&").Qual("golang-client/message/protoData", name+"s").Values(jen.Dict{
				jen.Id(name + "s"): jen.Id("protoList")}),
				jen.Id("stringList")),
		)
	}
	fmt.Printf("descp: %#v\n", descp)

}

func sysDataImplGen(impl *jen.File, name string, conf *YamlConfig) {
	impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("Default").Params(jen.Id("d").Qual("golang-client/bpcontext", "DobitInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
		jen.Comment("TODO implement me"),
		jen.Panic(jen.Lit("implement me")),
		jen.Comment("bs := BaseObject{fulldescription: \"\"}\n\t\treturn &BaseObjects{baseobjects: []*BaseObject{&bs}}"),
	)
	for _, desc := range conf.SystemData[name].Desc {
		if conf.Descriptor[desc] != nil {
			sortedDescriptorProperties := SortDescriptorProperty(conf.Descriptor[desc])
			for _, proerty := range sortedDescriptorProperties {
				impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id(proerty.Key).Params(jen.Id("d").Qual("golang-client/bpcontext", "DobitInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
					jen.Comment("TODO implement me"),
					jen.Panic(jen.Lit("implement me")),
				)
			}
		}
	}
	impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("SetServiceResponse").Params(jen.Id("index").Uint64(), jen.Id("response").Id("[]byte"), jen.Id("entity").Qual("golang-client/bpcontext", "DobitInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Block(
		jen.Comment("TODO implement me"),
		jen.Panic(jen.Lit("implement me")),
		jen.Comment("\tswitch index {\n\tcase 3001: //MemoryDistillToLongTerm\n\t\tprotoData := protoData.MemoryLogs{}\n"+
			"\t\terr := proto.Unmarshal(response, &protoData)\n\t\tif err != nil {\n"+
			"\t\t\tlog.Fatal(\"MemoryLogs Props ByteStream Handled Error: \", err)\n\t\t}\n"+
			"\t\tfmt.Print(\"MemoryDistillToLongTerm..MemoryLogs:\", protoData)\n\tcase 4001: //SummarizeAgent\n"+
			"\t\tprotoData := protoData.MemoryLogs{}\n\t\terr := proto.Unmarshal(response, &protoData)\n\t\tif err != nil {\n"+
			"\t\t\tlog.Fatal(\"MemoryLogs Props ByteStream Handled Error: \", err)\n\t\t}\n"+
			"\t\tfmt.Print(\"SummarizeAgent..MemoryLogs:\", protoData)\n\n\t}"),
	)
	// fmt.Printf("impl: %#v\n", impl)
}
func sysDataMgrGen(mgr *jen.File, name string, conf *YamlConfig) {
	mgr.Type().Id(name + "Manager").Struct(
		jen.Id(name + "s").Id(name + "s"),
	)
	mgr.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("GetDescriptor").Params(jen.Id("index").Uint64(), jen.Id("d").Qual("golang-client/bpcontext", "DobitInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
		jen.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit("func"), jen.Lit("GetDescriptor")),
		jen.Switch(jen.Id("index")).BlockFunc(func(g *jen.Group) {
			g.Case(jen.Lit(0)).Block(
				jen.Return(jen.Id("m").Dot("Default").Params(jen.Id("d"), jen.Id("ctx"))),
			)
			for _, desc := range conf.SystemData[name].Desc {
				if conf.Descriptor[desc] != nil {
					sortedDescriptorProperties := SortDescriptorProperty(conf.Descriptor[desc])
					for _, property := range sortedDescriptorProperties {
						desc_name := property.Key
						desc_i := property.Value
						g.Case(jen.Lit(desc_i)).Block(
							jen.Return(jen.Id("m").Dot(desc_name).Params(jen.Id("d"), jen.Id("ctx"))),
						)
					}
				}
			}
			g.Default().Block(
				jen.Id("log").Dot("Errorf").Call(jen.Lit("No such Descriptor in "+name+" Mgr")),
				jen.Return(jen.Id("m").Dot("Default").Params(jen.Id("d"), jen.Id("ctx"))),
			)
		}),
	)
	mgr.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("GetProps").Params(jen.Id("list").Qual("golang-client/bpcontext", "DataPropertyInterface"),
		jen.Id("index").Uint64()).Params(jen.Id("[]byte"), jen.Id("string")).Block(
		jen.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit(name+"Manager"), jen.Lit("GetProps")),
		jen.Id("listStruct, ok").Op(":=").Id("list").Assert(jen.Op("*").Id(name+"s")),
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Id("log").Dot("Debugf").Call(jen.Lit("Conversion failed.GetProps List does not hold a *"+name)),
		),
		jen.Id("interfaceObj, stringObj").Op(":=").Id("listStruct").Dot("GetPropIndex").Call(jen.Id("index")),
		jen.Id("serializeObj, ok").Op(":=").Id("interfaceObj").Assert(jen.Op("*").Qual("golang-client/message/protoData", name+"s")),
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Id("log").Dot("Debugf").Call(jen.Lit("Conversion failed.GetProps Return does not hold a *"+name)),
		),
		jen.Id("byteStream, err").Op(":=").Qual("google.golang.org/protobuf/proto", "Marshal").Call(jen.Id("serializeObj")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("log").Dot("Errorf").Call(jen.Lit(name+"s Props ByteStream Handled Error: %v"), jen.Id("err")),
		),
		jen.Return(jen.Id("byteStream"), jen.Id("stringObj")),
	)
	//fmt.Printf("mgr: %#v\n", mgr)
	err := mgr.Save("./implementation/mgr_gen_" + strings.ToLower(name) + ".go")
	if err != nil {
		return
	}
}
