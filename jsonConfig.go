package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type block struct {
	Info   map[string]string
	Blocks map[string]block
}

//OpInfoParser is a tool to parse and store data of a .JSON file
type OpInfoParser struct {
	Blocks map[string]block
}

func getKeyFromBlock(m map[string]interface{}) []string {
	tmpList := []string{}
	for key := range m {
		tmpList = append(tmpList, key)
	}
	return tmpList
}

func dispBlock(depth int, b map[string]block) {
	for key := range b {
		for i := 1; i < depth; i++ {
			fmt.Print("   ")
		}
		fmt.Println(">", key)
		for k, info := range b[key].Info {
			for i := 0; i < depth; i++ {
				fmt.Print("   ")
			}
			fmt.Println(">", k, ":", info)
		}
		dispBlock(depth+1, b[key].Blocks)
	}
}

func digInJSON(mapBlock map[string]interface{}) block {
	listBlock := getKeyFromBlock(mapBlock)
	b := block{make(map[string]string), make(map[string]block)}
	for _, v := range listBlock {
		if reflect.TypeOf(mapBlock[v]) != reflect.TypeOf("") {
			b.Blocks[v] = digInJSON(mapBlock[v].(map[string]interface{}))
		} else {
			b.Info[v] = mapBlock[v].(string)
		}
	}
	return b
}

//Init all info of the OpInfoParser
func (infoConfig *OpInfoParser) Init(pathFileConfig string) {
	file, err := ioutil.ReadFile(pathFileConfig)
	Check(err)
	var data interface{}
	json.Unmarshal(file, &data)
	m := data.(map[string]interface{})
	b := digInJSON(m)
	infoConfig.Blocks = b.Blocks
}

//NewOpInfoParser create and init a new OpInfoParser
func NewOpInfoParser(pathFileConfig string) OpInfoParser {
	file, err := ioutil.ReadFile(pathFileConfig)
	Check(err)
	var data interface{}
	json.Unmarshal(file, &data)
	m := data.(map[string]interface{})
	b := digInJSON(m)
	return OpInfoParser{Blocks: b.Blocks}
}

//Check if error occur and leave the program
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
