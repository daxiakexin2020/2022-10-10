package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
)

type PMap struct {
}

var _ data.PMap = (*PMap)(nil)

func NewPMap() data.PMap {
	return &PMap{}
}

var EmptyPMap = model.PMap{}

func (pm *PMap) Create(pmap model.PMap) (model.PMap, error) {
	return EmptyPMap, nil
}
func (pm *PMap) List() []model.PMap {
	return nil
}
func (pm *PMap) FetchPMap(id string) (model.PMap, error) {
	return EmptyPMap, nil
}
