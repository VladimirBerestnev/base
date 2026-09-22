package api

import (
	"1-converter/3-bin/config"
	"fmt"
)

type ApiDb struct {
	url string
	key string
}

func NewApiDb(url string, cfg *config.Config) *ApiDb {
	return &ApiDb{url: url,
		key: cfg.Key}
}

func (db *ApiDb) WriteBins(data string) {

	fmt.Println("Запись успешна")
}

func (db *ApiDb) ReadBins() ([]byte, error) {

	return []byte{}, nil
}
