package tools

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func descriptorGen(conf *DataYamlConfig) {
	f := jen.NewFilePathName("implementation", "implementation")
	sortedDescriptor := SortDescriptor(conf.Descriptor)
	for _, descEntry := range sortedDescriptor {
		fmt.Printf("%s: %v\n", descEntry.Key, descEntry.Value)
		f.Type().Id(descEntry.Key).IndexFunc(func(g *jen.Group) {
			for s := range descEntry.Value {
				g.Id(s).Params(jen.Id("d").Qual("golang-client/bpcontext", "DobitInterface")).Id("interface{}")
			}
		})
	}
	fmt.Printf("f: %#v\n", f)
	err := f.Save("./implementation/data_desc_interface_gen.go")
	if err != nil {
		return
	}
}
func entityGen(conf *DataYamlConfig) {
	f := jen.NewFilePathName("implementation", "implementation")
	sortedSystemData := SortData(conf.PluralData)
	sortedExternalData := SortData(conf.SingularData)
	f.Func().Id("InitMgrComponent").Params().BlockFunc(func(g *jen.Group) {
		for _, data := range sortedSystemData {
			g.Qual("golang-client/bpcontext", "RegisterMgrComponent").Call(jen.Lit(data.Value.Index), jen.Op("&").Id(data.Key+"Manager").Values())
		}
		for _, data := range sortedExternalData {
			g.Qual("golang-client/bpcontext", "RegisterMgrComponent").Call(jen.Lit(data.Value.Index), jen.Op("&").Id(data.Key+"Manager").Values())
		}
	})

	//fmt.Printf("f: %#v\n", f)
	err := f.Save("./implementation/entity_manager_gen.go")
	if err != nil {
		return
	}
}
