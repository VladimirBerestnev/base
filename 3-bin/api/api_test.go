package api_test

import (
	"1-converter/3-bin/api"
	"1-converter/3-bin/config"
	"1-converter/3-bin/storage"
	"testing"
)

func TestCreateByFlag(t *testing.T) {
	cfg := config.NewConfigTest()
	binApiDb := storage.NewBinListDb(api.NewApiDb("https://api.jsonbin.io/v3/b", cfg))
	expected := 200
	flagMap := map[string]string{
		"file": "../my.json",
		"name": "test",
	}
	id, err := api.CreateByFlag(binApiDb, flagMap)
	if err != nil {
		t.Errorf("Ошибка создания записи")
	}
	errDelete := api.DeleteByID(id, *cfg)
	if errDelete != expected {
		t.Errorf("Ошибка удаления записи")
	}
}

func TestUpdateByID(t *testing.T) {
	cfg := config.NewConfigTest()
	binApiDb := storage.NewBinListDb(api.NewApiDb("https://api.jsonbin.io/v3/b", cfg))
	expected := 200
	flagMap := map[string]string{
		"file": "../my.json",
		"name": "test",
	}
	id, err := api.CreateByFlag(binApiDb, flagMap)
	if err != nil {
		t.Errorf("Ошибка создания записи")
	}

	flagMap["id"] = id
	result := api.UpdateByID(flagMap, *cfg)
	if result != expected {
		t.Errorf("Ошибка обновления записи")
	}
	errDelete := api.DeleteByID(id, *cfg)
	if errDelete != expected {
		t.Errorf("Ошибка удаления записи")
	}
}

func TestDeleteByID(t *testing.T) {
	cfg := config.NewConfigTest()
	binApiDb := storage.NewBinListDb(api.NewApiDb("https://api.jsonbin.io/v3/b", cfg))
	expected := 200
	flagMap := map[string]string{
		"file": "../my.json",
		"name": "test",
	}
	id, err := api.CreateByFlag(binApiDb, flagMap)
	if err != nil {
		t.Errorf("Ошибка создания записи")
	}
	result := api.DeleteByID(id, *cfg)
	if result != expected {
		t.Errorf("Ошибка удаления записи")
	}
}
