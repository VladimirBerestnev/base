package api

import (
	"1-converter/3-bin/config"
	"1-converter/3-bin/file"
	"1-converter/3-bin/list"
	"1-converter/3-bin/storage"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ApiDb struct {
	url string
	key string
}

func NewApiDb(url string, cfg *config.Config) *ApiDb {
	return &ApiDb{url: url,
		key: cfg.Key}
}

func (db *ApiDb) WriteBins(data []byte) string {

	req, err := http.NewRequest(http.MethodPost, "https://api.jsonbin.io/v3/b", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", db.key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
	}
	fmt.Println("Ответ сервера:", string(body))
	defer resp.Body.Close()

	fmt.Println("StatusCode:", resp.StatusCode)

	var responce Response
	errDecode := json.Unmarshal(body, &responce)
	if errDecode != nil {
		fmt.Println("Ошибка декодирования")
	}
	return responce.Metadata.Id
}

func (db *ApiDb) ReadBins(id string) ([]byte, error) {
	baseurl, err := url.Parse("https://api.jsonbin.io/v3/b/" + id)
	params := url.Values{}
	params.Add("meta", "false")
	baseurl.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, baseurl.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", db.key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
	}

	defer resp.Body.Close()
	return body, err
}

func UpdateByID(mapFlags map[string]string, cfg config.Config) int {

	checkFile := file.CheckJsonFile(mapFlags["file"])
	if !checkFile {
		panic("File not .json")
	}
	baseurl, err := url.Parse("https://api.jsonbin.io/v3/b/" + mapFlags["id"])
	data, err := file.ReadAnyFile(mapFlags["file"])
	if err != nil {
		panic("Error read file")
	}

	req, err := http.NewRequest(http.MethodPut, baseurl.String(), bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", cfg.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
	}
	fmt.Println("Ответ сервера:", string(body))
	defer resp.Body.Close()
	return resp.StatusCode
}

func DeleteByID(id string, cfg config.Config) int {

	baseurl, err := url.Parse("https://api.jsonbin.io/v3/b/" + id)

	req, err := http.NewRequest(http.MethodDelete, baseurl.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", cfg.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
	}
	fmt.Println("Ответ сервера:", string(body))
	defer resp.Body.Close()
	list.DeleteList(id)
	return resp.StatusCode
}
func CreateByFlag(binApiDb *storage.BinListDb, mapFlags map[string]string) error {
	checkFile := file.CheckJsonFile(mapFlags["file"])
	if !checkFile {
		panic("File not .json")
	}
	fmt.Println(mapFlags["file"])
	data, err := file.ReadAnyFile(mapFlags["file"])
	if err != nil {
		panic("Error read file")
	}
	id := binApiDb.Db.WriteBins(data)

	NameID := list.NewName(mapFlags["name"], id)
	dataList := list.GetList()
	dataList.NameList = append(dataList.NameList, *NameID)
	data1, errMarsh := json.Marshal(dataList)
	if errMarsh != nil {
		return errors.New("Error Marshal")
	}
	errWrite := list.WriteList(data1)
	if errWrite != nil {
		return errors.New("Ошибка записи в файл")
	}
	return nil
}
