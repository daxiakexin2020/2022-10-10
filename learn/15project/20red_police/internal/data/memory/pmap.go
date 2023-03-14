package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"errors"
	"sync"
)

type PMap struct{}

type pmaps struct {
	list map[string]*model.PMap
	mu   sync.RWMutex
}

var _ data.PMap = (*PMap)(nil)
var (
	gpmaps    *pmaps
	EmptyPMap model.PMap
)

func init() {
	gpmaps = &pmaps{
		list: map[string]*model.PMap{},
	}
}

func NewPMap() data.PMap {
	return &PMap{}
}

func (pm *PMap) Create(pmap *model.PMap) (model.PMap, error) {
	gpmaps.mu.Lock()
	defer gpmaps.mu.Unlock()
	if _, ok := gpmaps.list[pmap.Id]; ok {
		return EmptyPMap, errors.New("地图已经存在")
	}
	gpmaps.list[pmap.Id] = pmap
	return *gpmaps.list[pmap.Id], nil
}
func (pm *PMap) List() []model.PMap {
	var res []model.PMap
	for _, pmap := range gpmaps.list {
		res = append(res, *pmap)
	}
	return res
}
func (pm *PMap) FetchPMap(id string) (model.PMap, error) {
	if pmap, ok := gpmaps.list[id]; ok {
		return *pmap, nil
	}
	return EmptyPMap, errors.New("不存在此地图")
}
