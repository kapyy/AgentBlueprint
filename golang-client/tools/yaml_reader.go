package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v2"
)

type DataFmt struct {
	Index    int `yaml:"index"`
	Property map[string]struct {
		Index int    `yaml:"index"`
		Type  string `yaml:"type"`
	} `yaml:"property"`
	Desc []string `yaml:"descriptor"`
}

type DetailedProperty struct {
	Index int
	Type  string
}
type DescriptorEntry struct {
	Key   string
	Value map[string]int
}
type KeyValue struct {
	Key   string
	Value int
}

type DataYamlConfig struct {
	Descriptor    map[string]map[string]int `yaml:"DataDescriptor"`
	SystemData    map[string]DataFmt        `yaml:"PluralDataIndex"`
	ExternalData  map[string]DataFmt        `yaml:"SingularDataIndex"`
	InternalData  map[string]DataFmt        `yaml:"InternalDataIndex"`
	ConnectorData map[string]DataFmt        `yaml:"ConnectionDataIndex"`
}

func (yc *DataYamlConfig) readYamlFile(filename string) *DataYamlConfig {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Yaml File Read Error: ", err)
	}
	// mf := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlFile, yc)
	if err != nil {
		log.Fatal("Yaml File Unmarshal Error: ", err)
	}
	return yc
}
func (yc *DataYamlConfig) getDataFromIndex(id int) (string, *DataEntry) {
	for key, value := range yc.SystemData {
		if value.Index == id {
			return "SystemData", &DataEntry{key, value}
		}
	}
	for key, value := range yc.ExternalData {
		if value.Index == id {
			return "ExternalData", &DataEntry{key, value}
		}
	}
	for key, value := range yc.InternalData {
		if value.Index == id {
			return "InternalData", &DataEntry{key, value}
		}
	}
	for key, value := range yc.ConnectorData {
		if value.Index == id {
			return "ConnectorData", &DataEntry{key, value}
		}
	}
	return "", nil

}

type FunctionFmt struct {
	Type     string `yaml:"Type"`
	ID       int    `yaml:"FunctionID"`
	InputID  int    `yaml:"InputNode"`
	OutputID int    `yaml:"OutputNode"`
}
type FunctionYamlConfig struct {
	Functions map[string]FunctionFmt `yaml:"Functions"`
}

func (fyc *FunctionYamlConfig) readYamlFile(filename string) *FunctionYamlConfig {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Yaml File Read Error: ", err)
	}
	err = yaml.Unmarshal(yamlFile, fyc)
	if err != nil {
		log.Fatal("Yaml File Unmarshal Error: ", err)
	}
	return fyc
}

type DataEntry struct {
	Key   string
	Value DataFmt
}

type sortedDataProperty struct {
	Key   string
	Value DetailedProperty
}
type sortedFunction struct {
	Key   string
	Value FunctionFmt
}

func SortDescriptor(descriptor map[string]map[string]int) []DescriptorEntry {
	var sortedDescriptor []DescriptorEntry
	for key, value := range descriptor {
		sortedDescriptor = append(sortedDescriptor, DescriptorEntry{key, value})
	}
	sort.Slice(sortedDescriptor, func(i, j int) bool {
		return sortedDescriptor[i].Key < sortedDescriptor[j].Key
	})
	return sortedDescriptor
}

func SortDescriptorProperty(sortedDataDescriptor map[string]int) []KeyValue {
	sortedValues := make([]KeyValue, 0, len(sortedDataDescriptor))
	for k, v := range sortedDataDescriptor {
		sortedValues = append(sortedValues, KeyValue{k, v})
	}
	sort.Slice(sortedValues, func(i, j int) bool {
		return sortedValues[i].Value < sortedValues[j].Value
	})
	return sortedValues
}

func SortData(data map[string]DataFmt) []DataEntry {
	var sortedData []DataEntry
	for key, value := range data {
		sortedData = append(sortedData, DataEntry{key, value})
	}
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i].Key < sortedData[j].Key
	})
	return sortedData
}

func SortDataProerty(data DataFmt) []sortedDataProperty {
	data_properties := data.Property
	var sorted_property []sortedDataProperty
	for key, detail_property := range data_properties {
		sorted_property = append(sorted_property, sortedDataProperty{Key: key, Value: DetailedProperty{Index: detail_property.Index, Type: detail_property.Type}})
	}
	sort.Slice(sorted_property, func(i, j int) bool { return sorted_property[i].Key < sorted_property[j].Key })
	return sorted_property
}
func SortFunction(functions map[string]FunctionFmt) []sortedFunction {
	var sortedFunctions []sortedFunction
	for key, value := range functions {
		sortedFunctions = append(sortedFunctions, sortedFunction{key, value})
	}
	sort.Slice(sortedFunctions, func(i, j int) bool {
		return sortedFunctions[i].Value.ID < sortedFunctions[j].Value.ID
	})
	return sortedFunctions
}
func ReadDataDescriptor() *DataYamlConfig {
	var dyc DataYamlConfig
	abspath, err := filepath.Abs("./config/DataIndexGen.yaml")
	if err != nil {
		log.Fatal("Yaml File Read Error: ", err)
	}
	conf := dyc.readYamlFile(abspath)
	fmt.Println("conf:", conf)

	return conf
}
func GenData(config *DataYamlConfig, overwrite bool) {
	descriptorGen(config)
	entityGen(config)
	internalDataGen(config)
	sysDataGen(config, overwrite) // Has overwrite para
	extDataGen(config, overwrite) // has overwrite para
	protoDataGen(config)
}
func ReadFunctions() *FunctionYamlConfig {
	var fyc FunctionYamlConfig
	abspath, err := filepath.Abs("./config/FunctionIndexGen.yaml")
	if err != nil {
		log.Fatal("Yaml File Read Error: ", err)
	}
	conf := fyc.readYamlFile(abspath)
	fmt.Println("conf:", conf)

	return conf
}
func GenFunction(dataConf *DataYamlConfig, functionConf *FunctionYamlConfig) {
	entityFunctionGen(functionConf, dataConf)
	entityInterfaceGen(functionConf, dataConf)
	reflectionGen(functionConf)
	protoFunctionGen(dataConf, functionConf)
}
