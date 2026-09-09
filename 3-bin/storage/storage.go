package storage

import (
	bin "3-bin/3-bin/bins"
	"encoding/json"
	"fmt"
	"os"
)

func SaveBins(binList bin.BinList, filename string) {
	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка сериализации json")
		return
	}
	file, errCreate := os.Create(filename)
	if errCreate != nil {
		fmt.Println("Ошибка создания файла")
		return
	}
	defer file.Close()

	_, errWrite := file.WriteString(string(data))
	if errWrite != nil {
		fmt.Println("Ошибка записи в файл")
		return
	}
	fmt.Println("Запись успешна")
}

func ReadBins(filename string) (bin.BinList, error) {
	bins := bin.NewBinList()
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return *bins, err
	}
	err = json.Unmarshal(data, &bins)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON")
		return *bins, err
	}
	return *bins, nil
}
