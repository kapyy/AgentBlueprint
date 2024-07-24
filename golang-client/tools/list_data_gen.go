package tools

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

func sysDataGen(conf *DataYamlConfig, overwrite bool) {
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

//func sysDataGen(conf *DataYamlConfig, overwrite bool) {
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
		g.Op("*").Qual("golang-client/message/proto", name)
	})
	descp.Func().Params(jen.Id("s").Id("*"+dataEntry.Key)).Id("Default").Call().Id("*").Qual("golang-client/message/proto", name).Block(
		//jen.Return(jen.Op("&").Qual("golang-client/message/proto", name).Values(jen.DictFunc(func(d jen.Dict) {
		//	for _, property := range sortedProperties {
		//		d[jen.Id(property.Key)] = jen.Id("s").Dot(property.Key).Params()
		//	}
		//}))),
		jen.Return(jen.Id("s").Dot(name)),
	)
	descp.Func().Params(jen.Id("s").Id("*" + dataEntry.Key)).Id("Set").Params(jen.Id(strings.ToLower(name)).Id("*").Qual("golang-client/message/proto", name)).BlockFunc(func(g *jen.Group) {
		for _, property := range sortedProperties {
			if isEmbeddedProperty(property.Value.Type) {
				g.Id("s").Dot("embedded_" + strings.ToLower(property.Key)).Dot("Set").Params(jen.Id(strings.ToLower(name)).Dot(property.Key))
			}
		}
		g.Id("s").Dot(name).Op("=").Id(strings.ToLower(name))
	})
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
		descp.Func().Params(jen.Id("s").Id("*" + name)).Id(prop_name + "String").Call().Id("string").BlockFunc(func(g *jen.Group) {
			g.Comment("Modify: this is where you define how you want your data to be recognized as natural language")
			//return fmt.Sprintf("Action's Description is: %v\n", s.ActionDescription())
			if isEmbeddedProperty(property.Value.Type) {
				g.Return(jen.Qual("fmt", "Sprintf").Call(jen.Lit(name+"'s "+prop_name+" is: %v\n"), jen.Id("s").Dot(prop_name).Call().Dot("FullString").Call()))
			} else {
				g.Return(jen.Qual("fmt", "Sprintf").Call(jen.Lit(name+"'s "+prop_name+" is: %v\n"), jen.Id("s").Dot(prop_name).Call()))
			}
		})
	}

	descp.Func().Params(jen.Id("s").Id("*"+name)).Id("Marshal").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		jen.Return(jen.Qual("google.golang.org/protobuf/proto", "Marshal").Call(jen.Id("s").Dot(name))),
	)
	descp.Func().Params(jen.Id("s").Id("*"+name)).Id("GetPropIndex").Params(jen.Id("index").Uint64()).Params(jen.Qual("golang-client/bpcontext", "DataPropertyInterface"), jen.Id("string")).Block(
		jen.Switch(jen.Id("index")).BlockFunc(func(g *jen.Group) {
			g.Case(jen.Id("0")).Block(
				jen.Return(jen.Id("s"), jen.Id("s").Dot("FullString").Call()),
			)
			for _, property := range sortedProperties {
				prop_name := property.Key
				prop := property.Value
				g.Case(jen.Id(fmt.Sprintf("%d", prop.Index))).Block(jen.Return(jen.Id("s").Id(",").Id("s").Dot(prop_name + "String").Call()))
			}
			g.Default().Block(
				jen.Return(jen.Id("s"), jen.Lit("")),
			)
		}),
	)

	if hasPlural {
		descp.Type().Id(name + "List").Struct(
			jen.Id(strings.ToLower(name + "List")).Id("[]*").Id(name),
		)
		descp.Func().Params(jen.Id("sl").Id("*" + name + "List")).Id("Set").Params(jen.Id(strings.ToLower(name+"List")).Id("*").Qual("golang-client/message/proto", name+"List")).BlockFunc(func(g *jen.Group) {
			g.Id("sl").Dot(strings.ToLower(name+"List")).Op("=").Make(jen.Index().Op("*").Id(name), jen.Id("0"))
			g.For(jen.Id("_,").Id("proto" + name).Op(":=").Range().Id(strings.ToLower(name + "List")).Dot(name + "List")).BlockFunc(func(gi *jen.Group) {
				gi.Id(strings.ToLower(name)).Op(":=").Op("&").Id(name).Values()
				gi.Id(strings.ToLower(name)).Dot("Set").Call(jen.Id("proto" + name))
				gi.Id("sl").Dot(strings.ToLower(name+"List")).Op("=").Append(jen.Id("sl").Dot(strings.ToLower(name+"List")), jen.Id(strings.ToLower(name)))
			})
		})
		//func (sl *ActionList) Marshal() ([]byte, error) {
		//	actionList := &proto.ActionList{}
		//	for _, action := range sl.actionlist {
		//		actionList.ActionList = append(actionList.ActionList, action.Action)
		//	}
		//	return proto.Marshal(actionList)
		//}
		descp.Func().Params(jen.Id("sl").Id("*"+name+"List")).Id("Marshal").Params().Params(jen.Index().Byte(), jen.Error()).Block(
			jen.Id(strings.ToLower(name)+"List").Op(":=").Op("&").Qual("golang-client/message/proto", name+"List").Values(),
			jen.For(jen.Id("_,").Id(strings.ToLower(name)).Op(":=").Range().Id("sl").Dot(strings.ToLower(name+"List"))).Block(
				jen.Id(strings.ToLower(name)+"List").Dot(name+"List").Op("=").Append(jen.Id(strings.ToLower(name)+"List").Dot(name+"List"), jen.Id(strings.ToLower(name)).Dot(name)),
			),
			jen.Return(jen.Qual("google.golang.org/protobuf/proto", "Marshal").Call(jen.Id(strings.ToLower(name)+"List"))),
		)

		descp.Func().Params(jen.Id("sl").Id("*"+name+"List")).Id("GetPropIndex").Params(jen.Id("index").Uint64()).Params(jen.Qual("golang-client/bpcontext", "DataPropertyInterface"), jen.Id("string")).Block(
			jen.Id("stringList").Op(":=").Lit(""),
			jen.For(jen.Id("i,").Id("s").Op(":=").Range().Id("sl").Dot(strings.ToLower(name+"List"))).Block(
				jen.Id("_").Op(",").Id("stringObj").Op(":=").Id("s").Dot("GetPropIndex").Call(jen.Id("index")),
				jen.Id("stringList").Op("+=").Qual("strconv", "Itoa").Call(jen.Id("i")).Op("+").Lit(". ").Op("+").Id("stringObj").Op("+").Lit("\n"),
			),
			jen.Return(jen.Id("sl"), jen.Id("stringList")),
		)
	}
	fmt.Printf("descp: %#v\n", descp)

}

func sysDataImplGen(impl *jen.File, name string, conf *DataYamlConfig) {
	impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("Default").Params(jen.Id("d").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
		jen.Comment("TODO implement me, this is where you read this data from, could be connected to a database or a service"),
		jen.Panic(jen.Lit("implement me")),
		jen.Comment("actions := &ActionList{}\n"+"actionList.Set(&proto.ActionList{\n"+"Actions:[]*proto.Action{\n"+"{\n"+"ActionDescription: \"go for a walk\",\n"+"Duration: 0,\n"+"StartTime: 0,\n"+"EndTime: 0,\n"+"},\n"+"},\n"+"})\n"+"return actionList\n"+"//ForPlural\n"+"action:=&Action{}\n"+"action.Set(&proto.Action{\n"+"ActionDescription: \"\",\n"+"Duration: 0,\n"+"StartTime: 0,\n"+"EndTime: 0,\n"+"})\n"+"return action"),
	)
	for _, desc := range conf.SystemData[name].Desc {
		if conf.Descriptor[desc] != nil {
			sortedDescriptorProperties := SortDescriptorProperty(conf.Descriptor[desc])
			for _, proerty := range sortedDescriptorProperties {
				impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id(proerty.Key).Params(jen.Id("d").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
					jen.Comment("TODO implement me, this is where you read this data from, could be connected to a database or a service"),
					jen.Panic(jen.Lit("implement me")),
				)
			}
		}
	}

	impl.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("SetServiceResponse").Params(jen.Id("index").Uint64(), jen.Id("response").Id("[]byte"), jen.Id("entity").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).BlockFunc(func(g *jen.Group) {
		g.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit(name+"Manager"), jen.Lit("SetServiceResponse"))
		g.Id("proto"+name+"List").Op(":=").Op("&").Qual("golang-client/message/proto", name+"List").Values()
		g.Id("err").Op(":=").Qual("google.golang.org/protobuf/proto", "Unmarshal").Call(jen.Id("response"), jen.Id("proto"+name+"List"))
		g.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("log").Dot("Errorf").Call(jen.Lit(name+"s Props ByteStream Handled Error: %s"), jen.Id("err")),
		)
		g.Id(strings.ToLower(name) + "List").Op(":=").Op("&").Id(name + "List").Values()
		g.Id(strings.ToLower(name) + "List").Dot("Set").Call(jen.Id("proto" + name + "List"))
		g.Id("ctx").Dot("SetResultData").Call(jen.Id(strings.ToLower(name) + "List"))
	})
	// fmt.Printf("impl: %#v\n", impl)
}
func sysDataMgrGen(mgr *jen.File, name string, conf *DataYamlConfig) {
	mgr.Type().Id(name + "Manager").Struct(
		jen.Id(name + "List").Id(name + "List"),
	)
	mgr.Func().Params(jen.Id("m").Id("*"+name+"Manager")).Id("GetDescriptor").Params(jen.Id("index").Uint64(), jen.Id("d").Qual("golang-client/bpcontext", "AgentInterface"), jen.Id("ctx").Qual("golang-client/bpcontext", "QueryContextInterface")).Qual("golang-client/bpcontext", "DataPropertyInterface").Block(
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
		jen.Id("listStruct, ok").Op(":=").Id("list").Assert(jen.Op("*").Id(name+"List")),
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Id("log").Dot("Debugf").Call(jen.Lit("Conversion failed.GetProps List does not hold a *"+name)),
		),
		jen.Id("interfaceObj, stringObj").Op(":=").Id("listStruct").Dot("GetPropIndex").Call(jen.Id("index")),
		jen.Id("serializeObj, ok").Op(":=").Id("interfaceObj").Assert(jen.Op("*").Qual("golang-client/message/proto", name+"List")),
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
