package tools

import (
	"math"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
func isBaseType(propType string) bool {
	return propType == "int32" || propType == "string" || propType == "bool" || propType == "uint64" || propType == "float32" || propType == "uint32"
}
func isEmbeddedProperty(prop_type string) bool {
	prop_name_noprefix, _ := strings.CutPrefix(prop_type, "[]")
	return !isBaseType(prop_name_noprefix)
}
func convertIfReference(prop_type string) string {

	prop_name_noprefix, isList := strings.CutPrefix(prop_type, "[]")
	propName := prop_name_noprefix
	if !isBaseType(prop_name_noprefix) {
		propName = "*" + propName
	}
	if isList {
		propName = "[]" + propName
	}
	return propName

}
func convertIfReferenceWithGrpc(prop_type string) string {

	prop_name_noprefix, isList := strings.CutPrefix(prop_type, "[]")
	propName := prop_name_noprefix
	if !isBaseType(prop_name_noprefix) {
		propName = "*protodata." + propName
	}
	if isList {
		propName = "[]" + propName
	}
	return propName
}
func convertIfPointer(prop_type string) string {

	prop_name_noprefix, isList := strings.CutPrefix(prop_type, "[]")
	propName := prop_name_noprefix
	if isList {
		return propName
	}
	if !isBaseType(prop_name_noprefix) {
		propName = "&" + propName
	}

	return propName
}

func convertProtoPropPtr(prop_type string) *jen.Statement {
	type_statement := jen.Id("")

	prop_name_noprefix, isList := strings.CutPrefix(prop_type, "[]")
	if isList {
		type_statement = jen.Id("[]")
	}
	if !isBaseType(prop_name_noprefix) {
		type_statement = type_statement.Id("*" + prop_name_noprefix)
	} else {
		type_statement = type_statement.Id(prop_name_noprefix)
	}
	return type_statement
}

func convertProtoPropObj(prop_type string) *jen.Statement {

	var type_statement *jen.Statement

	if prop_type == "int32" || prop_type == "uint64" || prop_type == "float32" || prop_type == "uint32" {
		type_statement = jen.Lit(0)
	} else if prop_type == "string" {
		type_statement = jen.Lit("")
	} else if prop_type == "bool" {
		type_statement = jen.Lit(false)
	} else {
		prop_name_noprefix, isList := strings.CutPrefix(prop_type, "[]")
		if isList {
			type_statement = jen.Id("[]")
			if prop_name_noprefix == "int32" {
				type_statement = type_statement.Id("int32").Values()
			} else if prop_name_noprefix == "uint64" {
				type_statement = type_statement.Id("uint64").Values()
			} else if prop_name_noprefix == "string" {
				type_statement = type_statement.Id("string").Values()
			} else if prop_name_noprefix == "float32" {
				type_statement = type_statement.Id("float32").Values()
			} else if prop_name_noprefix == "uint32" {
				type_statement = type_statement.Id("uint32").Values()
			} else if prop_name_noprefix == "bool" {
				type_statement = type_statement.Id("bool").Values()
			} else {
				type_statement = type_statement.Id("*").Qual("golang-client/message/protoData", prop_name_noprefix).Values()
			}

		} else {
			type_statement = jen.Id("&").Qual("golang-client/message/protoData", prop_name_noprefix).Values()
		}
	}
	return type_statement
}
func GetFunctionFullIndex(functionType string, functionIndex int) int {
	switch functionType {
	case "DefaultFunction":
		return int(1*math.Pow10(8)) + functionIndex
	case "MinorFunction":
		return int(2*math.Pow10(8)) + functionIndex
	case "StaticFunction":
		return int(3*math.Pow10(8)) + functionIndex
	default:
		return 0
	}
}
