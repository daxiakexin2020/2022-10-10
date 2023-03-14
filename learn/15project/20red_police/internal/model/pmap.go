package model

import "20red_police/tools"

type PMap struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

const (
	count_min = 1
	count_max = 8
)

func NewPMap(name string, count int) *PMap {
	if count == 0 {
		count = count_min
	}
	if count > 8 {
		count = count_max
	}
	return &PMap{
		Id:    tools.UUID(),
		Name:  name,
		Count: count,
	}
}
