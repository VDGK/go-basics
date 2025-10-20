package main

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, private bool, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New(`id cannot be empty`)
	}
	if name == "" {
		return nil, errors.New(`name cannot be empty`)
	}
	newBin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return newBin, nil
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}
