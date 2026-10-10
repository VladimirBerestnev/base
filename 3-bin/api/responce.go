package api

import "1-converter/3-bin/storage"

type Metadata struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Private   bool   `json:"private"`
}

type Response struct {
	Record   storage.BinList `json:"record"`
	Metadata Metadata        `json:"metadata"`
}
