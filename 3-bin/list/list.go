package list

import (
	"1-converter/3-bin/file"
	"1-converter/3-bin/storage"
	"encoding/json"
	"errors"
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

func GetListByID(binApiDb *storage.BinListDb, id string) error {
	data, errGet := binApiDb.Db.ReadBins(id)
	if errGet != nil {
		return errors.New("Read Error")
	}
	var list storage.BinList
	err := json.Unmarshal(data, &list)
	if err != nil {
		return errors.New("Unmarshal Error")
	}
	fmt.Println("Final:", list)
	return nil
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

func WriteList(data []byte) error {
	file, errCreate := os.Create("list.json")
	if errCreate != nil {
		return errors.New("Ошибка создания файла")
	}
	defer file.Close()

	_, errWrite := file.Write(data)
	if errWrite != nil {
		return errors.New("Ошибка записи в файл")
	}
	fmt.Println("Запись успешна")
	return nil
}
