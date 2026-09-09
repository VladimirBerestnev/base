package bin

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdat"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}
