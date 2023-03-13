package model

type PMap struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func NewPMap(name string, count int) *PMap {
	if count <= 0 {
		count = 1
	}
	return &PMap{Name: name, Count: count}
}
