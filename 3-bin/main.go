package main

import (
	"1-converter/3-bin/api"
	"1-converter/3-bin/config"
	"1-converter/3-bin/file"
	"1-converter/3-bin/storage"
	"fmt"
)

func main() {
	filename := "data.json"
	if file.CheckJsonFile(filename) {
		binListDb := storage.NewBinListDb(storage.NewJsonDb(filename))
		fmt.Println(binListDb.BinList)
	}

	cfg := config.NewConfig()

	binApiDb := storage.NewBinListDb(api.NewApiDb("http://ya.ru", cfg))
	fmt.Println(binApiDb.BinList)
}
