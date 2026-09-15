package main

import (
	"1-converter/3-bin/storage"
	"fmt"
)

func main() {
	binListDb := storage.NewBinListDb(storage.NewJsonDb("data.json"))
	fmt.Println(binListDb.BinList)
}
