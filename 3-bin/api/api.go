package api

import (
	"fmt"
)

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
