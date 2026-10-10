package main

import (
	"1-converter/3-bin/api"
	"1-converter/3-bin/config"
	"1-converter/3-bin/file"
	"1-converter/3-bin/list"
	"1-converter/3-bin/storage"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	ActivateApi()
}

func ActivateApi() {
	cfg := config.NewConfig()
	binApiDb := storage.NewBinListDb(api.NewApiDb("https://api.jsonbin.io/v3/b", cfg))

	mapFlags := CheckFlag()
	if mapFlags["create"] == "true" {
		CreateByFlag(binApiDb, mapFlags)
	}
	if mapFlags["update"] == "true" {
		api.UpdateByID(mapFlags, *cfg)
	}
	if mapFlags["delete"] == "true" {
		api.DeleteByID(mapFlags, *cfg)
	}
	if mapFlags["get"] == "true" {
		list.GetListByID(binApiDb, mapFlags)
	}
	if mapFlags["list"] == "true" {
		Namelist := list.GetList()
		for _, name := range Namelist.NameList {
			fmt.Println(name)
		}
	}
}

func CreateByFlag(binApiDb *storage.BinListDb, mapFlags map[string]string) {
	checkFile := file.CheckJsonFile(mapFlags["file"])
	if !checkFile {
		panic("File not .json")
	}
	fmt.Println(mapFlags["file"])
	data, err := file.ReadAnyFile(mapFlags["file"])
	if err != nil {
		panic("Error read file")
	}
	id := binApiDb.Db.WriteBins(data)

	NameID := list.NewName(mapFlags["name"], id)
	dataList := list.GetList()
	dataList.NameList = append(dataList.NameList, *NameID)
	data1, errMarsh := json.Marshal(dataList)
	if errMarsh != nil {
		fmt.Println("Error Marshal")
	}
	list.WriteList(data1)
}

func CheckFlag() map[string]string {
	file := flag.String("file", "", "Name of JSON-file")
	name := flag.String("name", "", "Name of bin-file")
	id := flag.String("id", "", "id of bin-file")
	create := flag.Bool("create", false, "Create bin-file")
	update := flag.Bool("update", false, "Update bin-file")
	delete := flag.Bool("delete", false, "Update bin-file")
	get := flag.Bool("get", false, "Get bin-file by id")
	list := flag.Bool("list", false, "List bin-files")
	flag.Parse()
	createString := strconv.FormatBool(*create)
	updateString := strconv.FormatBool(*update)
	deleteString := strconv.FormatBool(*delete)
	getString := strconv.FormatBool(*get)
	listString := strconv.FormatBool(*list)
	flagMap := map[string]string{
		"file":   *file,
		"name":   *name,
		"id":     *id,
		"create": createString,
		"update": updateString,
		"delete": deleteString,
		"get":    getString,
		"list":   listString,
	}
	return flagMap
}
