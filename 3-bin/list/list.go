package list

import (
	"1-converter/3-bin/file"
	"1-converter/3-bin/storage"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type NameID struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type NameList struct {
	NameList []NameID `json:"namelist"`
}

func NewName(name string, id string) *NameID {
	return &NameID{
		Id:   id,
		Name: name,
	}
}

func GetList() NameList {
	var NameList NameList
	data, err := file.ReadAnyFile("list.json")
	if err != nil {
		return NameList
	}
	errJson := json.Unmarshal(data, &NameList)
	if errJson != nil {
		return NameList
	}
	return NameList
}

func GetListByID(binApiDb *storage.BinListDb, mapFlags map[string]string) {
	data, errGet := binApiDb.Db.ReadBins(mapFlags["id"])
	if errGet != nil {
		fmt.Println("Read Error")
	}
	var list storage.BinList
	err := json.Unmarshal(data, &list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Final:", list)
}

func DeleteList(id string) {
	list := GetList()
	for i, name := range list.NameList {
		if name.Id == id {
			list.NameList = slices.Delete(list.NameList, i, i+1)
		}
	}
	file, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	WriteList(file)

}

func WriteList(data []byte) {
	file, errCreate := os.Create("list.json")
	if errCreate != nil {
		fmt.Println("Ошибка создания файла")
		return
	}
	defer file.Close()

	_, errWrite := file.Write(data)
	if errWrite != nil {
		fmt.Println("Ошибка записи в файл")
		return
	}
	fmt.Println("Запись успешна")
}
