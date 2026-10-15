package list_test

import (
	"1-converter/3-bin/api"
	"1-converter/3-bin/config"
	"1-converter/3-bin/list"
	"1-converter/3-bin/storage"
	"testing"
)

func TestGetListByID(t *testing.T) {
	cfg := config.NewConfigTest()
	expected := 200
	binApiDb := storage.NewBinListDb(api.NewApiDb("https://api.jsonbin.io/v3/b", cfg))
	flagMap := map[string]string{
		"file": "../my.json",
		"name": "test",
	}
	id, err := api.CreateByFlag(binApiDb, flagMap)
	if err != nil {
		t.Errorf("Ошибка создания записи")
	}
	result := list.GetListByID(binApiDb, id)
	if result != nil {
		t.Errorf("Ошибка чтения записи")
	}
	errDelete := api.DeleteByID(id, *cfg)
	if errDelete != expected {
		t.Errorf("Ошибка удаления записи")
	}
}
