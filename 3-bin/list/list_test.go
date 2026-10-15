package list_test

import (
	"1-converter/3-bin/api"
	"1-converter/3-bin/config"
	"1-converter/3-bin/list"
	"1-converter/3-bin/storage"
	"testing"
)

func TestGetListByID(t *testing.T) {
	cfg := config.NewConfig()
	binApiDb := storage.NewBinListDb(api.NewApiDb("https://api.jsonbin.io/v3/b", cfg))
	id := "6aca2c02ffd5d160535ee4db"
	result := list.GetListByID(binApiDb, id)
	if result != nil {
		t.Errorf("Ошибка чтения записи")
	}
}
