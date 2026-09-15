package storage

import (
	bin "3-bin/3-bin/bins"
	"encoding/json"
	"fmt"
	"os"
)

func SaveBins(binList bin.BinList, filename string) error {
	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка сериализации json")
		return err
	}
	file, errCreate := os.Create(filename)
	if errCreate != nil {
		fmt.Println("Ошибка создания файла")
		return err
	}
	defer file.Close()

	_, errWrite := file.WriteString(string(data))
	if errWrite != nil {
		fmt.Println("Ошибка записи в файл")
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}

func ReadBins(filename string) (bin.BinList, error) {
	var bins bin.BinList
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return bins, err
	}
	err = json.Unmarshal(data, &bins)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON")
		return bins, err
	}
	return bins, nil
}
