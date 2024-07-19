package tools

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func entityFunctionGen(conf *FunctionYamlConfig) {
	f := jen.NewFilePathName("client_object", "dobit_client")
	sortFunctions := SortFunction(conf.Functions)
	for _, function := range sortFunctions {
		switch function.Value.Type {
		case "DefaultFunction":
			f.Func().Params(jen.Id("d").Op("*").Id("AgentEntity")).Id(function.Key).Params(jen.Id("i").Op("*").Qual("golang-client/bpcontext", "DataInstance"), jen.Id("input").Qual("golang-client/bpcontext", "UserInput")).Block(
				jen.Id("queryctx").Op(":=").Id("i").Dot("LinkNewQueryContext").Call(jen.Id("i").Dot("Data").Call()),
				jen.Id("query").Op(":=").Id("d").Dot("Queries").Index(jen.Lit(GetFunctionFullIndex(function.Value.Type, function.Value.ID))),
				jen.Id("ctx").Op(":=").Qual("golang-client/bpcontext", "NewAgentContext").Call(jen.Id("d"), jen.Id("queryctx")),
				jen.Id("queryctx").Dot("SetInputText").Call(jen.Op("&").Id("input").Dot("InputText")),
				jen.Id("query").Dot("BPFunctionNodes").Dot("FunctionParam").Dot("InputText").Op("=").Op("&").Id("input").Dot("InputText"),
				jen.Id("query").Dot("call").Call(jen.Id("d"), jen.Id("ctx")),
			)
		case "StaticFunction":
			f.Func().Params(jen.Id("d").Op("*").Id("AgentEntity")).Id(function.Key).Params(jen.Id("i").Op("*").Qual("golang-client/bpcontext", "DataInstance")).Block(
				jen.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit("func"), jen.Lit(function.Key)),
				jen.Id("byteData").Op(",").Id("err").Op(":=").Id("i").Dot("Data").Call().Dot("Marshal").Call(),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id("log").Dot("Errorf").Call(jen.Lit("i.Data().Marshal error: %v"), jen.Id("err")),
					jen.Return(),
				),
				jen.Id("queryCtx").Op(":=").Id("i").Dot("LinkNewQueryContext").Call(jen.Id("i").Dot("Data").Call()),

				jen.Id("err").Op("=").Id("d").Dot("callSubordinateFunction").Call(jen.Lit(GetFunctionFullIndex(function.Value.Type, function.Value.ID)), jen.Lit(function.Value.InputID), jen.Id("byteData"), jen.Id("queryCtx")),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id("log").Dot("Errorf").Call(jen.Lit("callSubordinateFunction error: %v"), jen.Id("err")),
					jen.Return(),
				),
			)
		}
	}
	fmt.Printf("f:%#v\n", f)
	err := f.Save("client_object/entity_function_gen.go")
	if err != nil {
		return
	}
}

//type AgentInterface interface {
//
//	//Active PipelineFunctions
//	InsertActionsWithObservation(datactx QueryContextInterface)
//
//	GetDataManager(mgr int) DataManagerInterface
//}

func entityInterfaceGen(conf *FunctionYamlConfig) {
	f := jen.NewFilePathName("bpcontext", "bpcontext")
	f.Type().Id("AgentInterface").InterfaceFunc(func(g *jen.Group) {
		sortFunctions := SortFunction(conf.Functions)
		for _, function := range sortFunctions {
			switch function.Value.Type {
			case "DefaultFunction":
				g.Id(function.Key).Params(jen.Id("i").Op("*").Id("DataInstance"), jen.Id("input").Id("UserInput"))
			case "StaticFunction":
				g.Id(function.Key).Params(jen.Id("i").Op("*").Id("DataInstance"))
			}
		}
		g.Id("GetDataManager").Params(jen.Id("mgr").Int()).Id("DataManagerInterface")
	})

	fmt.Printf("f:%#v\n", f)
	err := f.Save("bpcontext/agent_interface_gen.go")
	if err != nil {
		return
	}
}
