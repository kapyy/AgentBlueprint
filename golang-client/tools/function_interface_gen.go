package tools

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
)

func entityFunctionGen(conf *FunctionYamlConfig, dataConf *DataYamlConfig) {
	f := jen.NewFilePathName("client_object", "dobit_client")
	sortFunctions := SortFunction(conf.Functions)
	for _, function := range sortFunctions {
		_, outData := dataConf.getDataFromIndex(function.Value.OutputID)
		switch function.Value.Type {
		case "DefaultFunction":
			var outDataName string
			if outDataName = outData.Key; dataConf.isPluralData(function.Value.OutputID) {
				outDataName = outData.Key + "List"
			}
			f.Func().Params(jen.Id("d").Op("*").Id("AgentEntity")).Id(function.Key).Params(jen.Id("i").Op("*").Qual("golang-client/bpcontext", "DataInstance"), jen.Id("input").Qual("golang-client/bpcontext", "UserInput")).Id("*").Qual("golang-client/implementation", outDataName).BlockFunc(func(g *jen.Group) {
				g.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit("func"), jen.Lit(function.Key))
				g.Id("queryctx").Op(":=").Id("i").Dot("LinkNewQueryContext").Call(jen.Id("i").Dot("Data").Call())
				g.Id("query").Op(":=").Id("d").Dot("Queries").Index(jen.Lit(GetFunctionFullIndex(function.Value.Type, function.Value.ID)))
				g.Id("ctx").Op(":=").Qual("golang-client/bpcontext", "NewAgentContext").Call(jen.Id("d"), jen.Id("queryctx"))
				g.Id("queryctx").Dot("SetInputText").Call(jen.Op("&").Id("input").Dot("InputText"))
				g.Id("query").Dot("BPFunctionNodes").Dot("FunctionParam").Dot("InputText").Op("=").Op("&").Id("input").Dot("InputText")
				g.Id("query").Dot("call").Call(jen.Id("d"), jen.Id("ctx"))
				g.Id(strings.ToLower(outDataName)).Op(",").Id("ok").Op(":=").Id("queryctx").Dot("ResultData").Call().Assert(jen.Op("*").Qual("golang-client/implementation", outDataName))
				g.If(jen.Op("!").Id("ok")).Block(jen.Id("log").Dot("Errorf").Call(jen.Lit(function.Key + ": Return data type error")))
				g.Return(jen.Id(strings.ToLower(outDataName)))
			})
		case "StaticFunction":
			var outDataName string
			if outDataName = outData.Key; dataConf.isPluralData(function.Value.OutputID) {
				outDataName = outData.Key + "List"
			}
			f.Func().Params(jen.Id("d").Op("*").Id("AgentEntity")).Id(function.Key).Params(jen.Id("i").Op("*").Qual("golang-client/bpcontext", "DataInstance")).Id("*").Qual("golang-client/implementation", outDataName).Block(
				jen.Id("log").Op(":=").Qual("golang-client/modules/logger", "GetLogger").Call().Dot("WithField").Call(jen.Lit("func"), jen.Lit(function.Key)),
				jen.Id("byteData").Op(",").Id("err").Op(":=").Id("i").Dot("Data").Call().Dot("Marshal").Call(),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id("log").Dot("Errorf").Call(jen.Lit("i.Data().Marshal error: %v"), jen.Id("err")),
					jen.Return(jen.Nil()),
				),
				jen.Id("queryctx").Op(":=").Id("i").Dot("LinkNewQueryContext").Call(jen.Id("i").Dot("Data").Call()),

				jen.Id("err").Op("=").Id("d").Dot("callSubordinateFunction").Call(jen.Lit(GetFunctionFullIndex(function.Value.Type, function.Value.ID)), jen.Lit(function.Value.InputID), jen.Id("byteData"), jen.Id("queryctx")),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id("log").Dot("Errorf").Call(jen.Lit("callSubordinateFunction error: %v"), jen.Id("err")),
					jen.Return(jen.Nil()),
				),
				jen.Id(strings.ToLower(outDataName)).Op(",").Id("ok").Op(":=").Id("queryctx").Dot("ResultData").Call().Assert(jen.Op("*").Qual("golang-client/implementation", outDataName)),
				jen.If(jen.Op("!").Id("ok")).Block(jen.Id("log").Dot("Errorf").Call(jen.Lit(function.Key+": Return data type error"))),
				jen.Return(jen.Id(strings.ToLower(outDataName))),
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

func entityInterfaceGen(conf *FunctionYamlConfig, dataConf *DataYamlConfig) {
	f := jen.NewFilePathName("bpcontext", "bpcontext")
	f.Type().Id("AgentInterface").InterfaceFunc(func(g *jen.Group) {
		sortFunctions := SortFunction(conf.Functions)
		for _, function := range sortFunctions {
			_, outData := dataConf.getDataFromIndex(function.Value.OutputID)
			var outDataName string
			if outDataName = outData.Key; dataConf.isPluralData(function.Value.OutputID) {
				outDataName = outData.Key + "List"
			}
			switch function.Value.Type {
			case "DefaultFunction":
				g.Id(function.Key).Params(jen.Id("i").Op("*").Id("DataInstance"), jen.Id("input").Id("UserInput")).Id("*").Qual("golang-client/implementation", outDataName)
			case "StaticFunction":
				g.Id(function.Key).Params(jen.Id("i").Op("*").Id("DataInstance")).Id("*").Qual("golang-client/implementation", outDataName)
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

//	func GetDataIDFromFunctionID(functionID uint64) uint64 {
//		switch functionID {
//		case 100100100:
//			return 1001 //SomeFunction>>Actions
//		default:
//			return 1
//		}
//	}
func reflectionGen(conf *FunctionYamlConfig) {
	f := jen.NewFilePathName("implementation", "implementation")
	f.Func().Id("GetDataIDFromFunctionID").Params(jen.Id("functionID").Uint64()).Uint64().Block(
		jen.Switch(jen.Id("functionID")).BlockFunc(func(g *jen.Group) {
			sortFunctions := SortFunction(conf.Functions)
			for _, function := range sortFunctions {
				if function.Value.Type != "MinorFunction" {
					g.Case(jen.Lit(GetFunctionFullIndex(function.Value.Type, function.Value.ID))).Block(jen.Return(jen.Lit(function.Value.OutputID)))
				}

			}
			g.Default().Block(
				jen.Return(jen.Lit(1)),
			)
		}),
	)

	fmt.Printf("f:%#v\n", f)
	err := f.Save("implementation/service_reflection_gen.go")
	if err != nil {
		return
	}
}
