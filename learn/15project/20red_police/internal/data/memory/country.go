package memory

import (
	"20red_police/common"
	"20red_police/internal/data"
	"20red_police/internal/model"
	"errors"
	"sync"
)

const (
	SovietUnion   = "苏联"
	Cuba          = "古巴"
	Iraq          = "伊拉克"
	Syria         = "叙利亚"
	Libya         = "利比亚"
	United_States = "美国"
	Britain       = "英国"
	France        = "法国"
	Germany       = "德国"
	South_Korea   = "韩国"
)

type Country struct{}

func NewCountry() data.Country {
	c := &Country{}
	data.GclassTree().Register(c)
	return c
}

type gcountrysOuter struct {
	list []model.Country
	c    map[string]*model.Country
}

var _ data.Country = (*Country)(nil)

var (
	gcOuter      *gcountrysOuter
	gcOnce       sync.Once
	emptyCountry model.Country
)

var gSovietUnion = &model.Country{
	Id:                "1",
	Name:              SovietUnion,
	ArchitectureNames: []string{sa_base, sa_barrack},
}

var gCubd = &model.Country{
	Id:                "2",
	Name:              Cuba,
	ArchitectureNames: []string{sa_base, sa_barrack},
}

var gUS = &model.Country{
	Id:                "3",
	Name:              United_States,
	ArchitectureNames: []string{af_barrack},
}

func init() {
	gcOuter = &gcountrysOuter{c: map[string]*model.Country{}, list: make([]model.Country, 0)}
	gcOuter.c[SovietUnion] = gSovietUnion
	gcOuter.c[Cuba] = gCubd
	gcOuter.c[United_States] = gUS
	gcOuter.list = []model.Country{*gSovietUnion, *gCubd, *gUS}
}

func (c *Country) CountryList() []model.Country {
	return gcOuter.list
}

func (c *Country) FetchCountry(name string) (model.Country, error) {
	if country, ok := gcOuter.c[name]; ok {
		return *country, nil
	}
	return emptyCountry, errors.New("no this country")
}

func (c *Country) Name() string {
	return common.REGISTER_MEMORY_PMAP
}
