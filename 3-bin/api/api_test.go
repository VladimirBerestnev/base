package api_test

import (
	"1-converter/3-bin/api"
	"1-converter/3-bin/config"
	"1-converter/3-bin/storage"
	"testing"
)

func TestCreateByFlag(t *testing.T) {
	cfg := config.NewConfig()
	binApiDb := storage.NewBinListDb(api.NewApiDb("https://api.jsonbin.io/v3/b", cfg))

	flagMap := map[string]string{
		"file": "my.json",
		"name": "test",
	}
	err := api.CreateByFlag(binApiDb, flagMap)
	if err != nil {
		t.Errorf("Ошибка создания записи")
	}
}

func TestUpdateByID(t *testing.T) {
	cfg := config.NewConfig()
	expected := 200

	flagMap := map[string]string{
		"file": "my.json",
		"id":   "6aca2c02ffd5d160535ee4db",
	}
	result := api.UpdateByID(flagMap, *cfg)
	if result != expected {
		t.Errorf("Ошибка обновления записи")
	}
}

func TestDeleteByID(t *testing.T) {
	cfg := config.NewConfig()
	expected := 200
	id := "6aca2c02ffd5d160535ee4db"
	result := api.DeleteByID(id, *cfg)
	if result != expected {
		t.Errorf("Ошибка удаления записи")
	}
}
