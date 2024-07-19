package tools

import (
	"github.com/dave/jennifer/jen"
	"strings"
)

func internalDataGen(conf *DataYamlConfig) {
	f := jen.NewFilePathName("implementation", "implementation")
	sortedInternalData := SortData(conf.InternalData)
	for _, data := range sortedInternalData {
		name := data.Key
		f.Type().Id(name).StructFunc(func(g *jen.Group) {
			g.Op("*").Qual("golang-client/message/protoData", name)
		})
		f.Func().Params(jen.Id("s").Id("*"+name)).Id("Default").Call().Id("*").Qual("golang-client/message/protoData", name).Block(
			//jen.Comment("TODO: implement me"),
			//jen.Panic(jen.Lit("implement me")),
			jen.Return(jen.Id("s").Dot(name)),
		)
		f.Func().Params(jen.Id("s").Id("*" + name)).Id("Set").Params(jen.Id(strings.ToLower(name)).Id("*").Qual("golang-client/message/protoData", name)).Block(
			jen.Id("s").Dot(name).Op("=").Id(strings.ToLower(name)),
		)
		f.Func().Params(jen.Id("s").Id("*"+name)).Id("FullString").Call().Id("string").Block(
			jen.Comment("TODO: implement me"),
			jen.Panic(jen.Lit("implement me")),
		)
	}
	err := f.Save("./implementation/internal_data_gen.go")
	if err != nil {
		return
	}
}
