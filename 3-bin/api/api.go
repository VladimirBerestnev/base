package api

import (
	"1-converter/3-bin/config"
	"1-converter/3-bin/file"
	"1-converter/3-bin/list"
	"bytes"
	"encoding/json"
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

	fmt.Println("Запись успешна")

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
	return body, nil
}

func UpdateByID(mapFlags map[string]string, cfg config.Config) {

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
}

func DeleteByID(mapFlags map[string]string, cfg config.Config) {

	baseurl, err := url.Parse("https://api.jsonbin.io/v3/b/" + mapFlags["id"])

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
	list.DeleteList(mapFlags["id"])
}
