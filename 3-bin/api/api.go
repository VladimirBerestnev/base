package api

import (
	"1-converter/3-bin/config"
	"fmt"
)

func GetConfig() {
	config := config.NewConfig()
	fmt.Println(config.Key)
}

type ApiDb struct {
	url string
}

func NewApiDb(url string) *ApiDb {
	return &ApiDb{url: url}
}

func (db *ApiDb) WriteBins(data string) {

	fmt.Println("Запись успешна")
}

func (db *ApiDb) ReadBins() ([]byte, error) {

	return []byte{}, nil
}
