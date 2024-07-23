package tools

import (
	"fmt"
	"os"
)

func protoDataGen(conf *DataYamlConfig) {
	f, err := os.Create("../message/proto/dataIndexGen.proto")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString("syntax = \"proto3\";\n")
	f.WriteString("package protoData;\n")
	f.WriteString("option go_package = \"golang-client/message\";\n")
	f.WriteString("// --------------InternalData--------\n")
	sortedInternalData := SortData(conf.InternalData)
	for _, data := range sortedInternalData {
		name := data.Key
		f.WriteString("message " + name + " {\n")
		for propname, detail_property := range data.Value.Property {
			f.WriteString("	" + detail_property.Type + " " + ToSnakeCase(propname) + " = " + fmt.Sprint(detail_property.Index) + ";\n")
		}
		f.WriteString("}\n")
	}
	f.WriteString("// --------------SystemData--------\n")
	sortedSystemData := SortData(conf.SystemData)
	for _, data := range sortedSystemData {
		name := data.Key
		f.WriteString("message " + name + " {\n")
		for propname, detail_property := range data.Value.Property {
			f.WriteString("	" + detail_property.Type + " " + ToSnakeCase(propname) + " = " + fmt.Sprint(detail_property.Index) + ";\n")
		}
		f.WriteString("}\n")
		f.WriteString("message " + name + "List {\n")
		f.WriteString("	repeated " + name + " " + ToSnakeCase(name) + "_list" + " = 1;\n")
		f.WriteString("}\n")
	}
	f.WriteString("// --------------ExternalData--------\n")
	sortedExternalData := SortData(conf.ExternalData)
	for _, data := range sortedExternalData {
		name := data.Key
		f.WriteString("message " + name + " {\n")
		for propname, detail_property := range data.Value.Property {
			f.WriteString("	" + detail_property.Type + " " + ToSnakeCase(propname) + " = " + fmt.Sprint(detail_property.Index) + ";\n")
		}
		f.WriteString("}\n")
	}
	f.Close()
}

//syntax = "proto3";
//package protoData;
//option go_package = "message/grpc";
//import "dataIndexGen.proto";
//
////Internal Python Service to distribute the apm request to individual functions
//service APMFunctionsService{
//
//
//rpc InsertActionsWithObservation(TextPyRequest) returns(Actions);
//rpc ActionFormatter(Action) returns (ParsedAction);
//
//}
//message GeneralPyRequest{
//string prompt = 1;
//string text = 2;
//optional string system_prompt = 2;

func protoFunctionGen(dataConf *DataYamlConfig, conf *FunctionYamlConfig) {
	f, err := os.Create("../message/proto/functionDistribute.proto")
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("syntax = \"proto3\";\n")
	f.WriteString("package proto;\n")
	f.WriteString("option go_package = \"golang-client/message\";\n")
	f.WriteString("import \"message/proto/dataIndexGen.proto\";\n")
	f.WriteString("//Internal Python Service to distribute the apm request to individual functions\n")
	f.WriteString("service APMFunctionsService{\n")
	sortFunction := SortFunction(conf.Functions)
	for _, function := range sortFunction {
		name := function.Key
		var inputData, outputData string
		switch function.Value.Type {
		case "DefaultFunction":
			inputData = "GeneralPyRequest"
			outDataType, outDataName := dataConf.getDataFromIndex(function.Value.OutputID)
			if outDataType == "SystemData" {
				outputData = outDataName.Key + "List"
			} else {
				outputData = outDataName.Key
			}
		case "StaticFunction":
			_, inDataName := dataConf.getDataFromIndex(function.Value.InputID)
			inputData = inDataName.Key
			_, outDataName := dataConf.getDataFromIndex(function.Value.OutputID)
			outputData = outDataName.Key
		default:
		}
		f.WriteString("rpc " + name + "(" + inputData + ") returns(" + outputData + ");\n")
	}
	f.WriteString("}\n")
	f.WriteString("message GeneralPyRequest{\n")
	f.WriteString("string prompt = 1;\n")
	f.WriteString("string text = 2;\n")
	f.WriteString("optional string system_prompt = 3;\n")
	f.WriteString("}\n")
	f.Close()
}
