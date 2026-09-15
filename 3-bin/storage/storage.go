package storage

import (
	bin "1-converter/3-bin/bins"
	"encoding/json"
	"fmt"
	"os"
)

type BinList struct {
	Bins []bin.Bin `json:"bins"`
}

type BinListDb struct {
	BinList
	db Db
}

func NewBinListDb(db Db) *BinListDb {
	file, err := db.ReadBins()
	if err != nil {
		return &BinListDb{
			BinList: BinList{
				Bins: []bin.Bin{},
			},
			db: db,
		}
	}
	var list BinList
	err = json.Unmarshal(file, &list)
	if err != nil {
		return &BinListDb{
			BinList: BinList{
				Bins: []bin.Bin{},
			},
			db: db,
		}
	}
	return &BinListDb{
		BinList: list,
		db:      db,
	}
}

type Db interface {
	WriteBins(string)
	ReadBins() ([]byte, error)
}

type JsonDb struct {
	filename string
}

func NewJsonDb(filename string) *JsonDb {
	return &JsonDb{filename: filename}
}

func (bins *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(bins)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (bins *BinListDb) SaveBins() error {
	data, err := bins.BinList.ToBytes()
	if err != nil {
		return err
	}
	bins.db.WriteBins(string(data))
	return nil
}

func (db *JsonDb) WriteBins(data string) {
	file, errCreate := os.Create(db.filename)
	if errCreate != nil {
		fmt.Println("Ошибка создания файла")
		return
	}
	defer file.Close()

	_, errWrite := file.WriteString(data)
	if errWrite != nil {
		fmt.Println("Ошибка записи в файл")
		return
	}
	fmt.Println("Запись успешна")
}

func (db *JsonDb) ReadBins() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}
	return data, nil
}
