package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"errors"
	"sync"
)

type Architecture struct{}

type aouter struct {
	list []model.Architecture
	as   map[string]*model.Architecture
}

type afBase struct {
	*model.Architecture
}

type saBase struct {
	*model.Architecture
}

type barracks struct {
	*model.Architecture
}

type afbarracks struct {
	*model.Architecture
	*barracks
}

type ukBarracks struct {
	*model.Architecture
	*afbarracks
}

type saBarracks struct {
	*model.Architecture
	*barracks
}

type iraqBarracks struct {
	*model.Architecture
	*saBarracks
}

var _ data.Architecture = (*Architecture)(nil)

var emptyArchitecture model.Architecture

var saba = &saBase{
	&model.Architecture{
		Id:                "00",
		Name:              sa_barrack,
		ArmList:           []string{},
		ConstructionPrice: 100,
		CanBeBuildNum:     1,
		CanPutAway:        true,
	},
}

var afba = &afBase{
	&model.Architecture{
		Id:            "01",
		Name:          af_barrack,
		ArmList:       []string{},
		CanBeBuildNum: 1,
		CanPutAway:    true,
	},
}

var b = &barracks{
	&model.Architecture{
		Id:      "1",
		Name:    barrack,
		ArmList: []string{engineer},
	},
}

var afb = &afbarracks{
	Architecture: &model.Architecture{
		Id:      "2",
		Name:    af_barrack,
		ArmList: []string{af_soldier, af_flying_soldier, af_spy, af_dog, af_chronosphere_soldier},
	},
	barracks: b,
}

var ukb = &ukBarracks{
	Architecture: &model.Architecture{
		Id:      "3",
		Name:    uk_barrack,
		ArmList: []string{british_sniper},
	},
	afbarracks: afb,
}

var sab = &saBarracks{
	Architecture: &model.Architecture{
		Id:                "4",
		Name:              sa_barrack,
		ArmList:           []string{sa_soldier, sa_dog},
		BloodVolume:       C_BloodVolume_100,
		ConstructionPrice: C_ConstructionPrice_500,
	},
	barracks: b,
}

var iraqb = &iraqBarracks{
	Architecture: &model.Architecture{
		Id:      "5",
		Name:    iraq_barrack,
		ArmList: []string{iraq_badiation_engineer},
	},
	saBarracks: sab,
}

var (
	aronce  sync.Once
	gaouter *aouter
)

const (
	sa_base      = "苏军基地"
	af_base      = "盟军基地"
	barrack      = "兵营"
	af_barrack   = "盟军兵营"
	uk_barrack   = "英国兵营"
	sa_barrack   = "sjby"
	iraq_barrack = "伊拉克兵营"
)

const (
	C_BloodVolume_100 = 100
	C_BloodVolume_200 = 200
	C_BloodVolume_300 = 300
	C_BloodVolume_400 = 400
	C_BloodVolume_500 = 500
	C_BloodVolume_600 = 600
)

const (
	C_ConstructionPrice_100 = 100
	C_ConstructionPrice_200 = 200
	C_ConstructionPrice_300 = 300
	C_ConstructionPrice_400 = 400
	C_ConstructionPrice_500 = 500
	C_ConstructionPrice_600 = 600
	C_ConstructionPrice_700 = 700
	C_ConstructionPrice_800 = 800
)

func init() {
	gaouter = defaultAouter()
	aggregationAFBarracks()
	aggregationUKBarracks()
	aggregationSABarracks()
	aggregationIRAQBarracks()
	arlist()
}

func NewArchitecture() data.Architecture {
	return &Architecture{}
}

func aggregationAFBarracks() *model.Architecture {
	a := &model.Architecture{
		Id:          afb.Id,
		Name:        afb.Name,
		ArmList:     append(afb.ArmList, afb.barracks.ArmList...),
		BloodVolume: afb.BloodVolume,
	}
	gaouter.as[afb.Name] = a
	return a
}

func aggregationUKBarracks() *model.Architecture {
	ukb.ArmList = append(ukb.ArmList, aggregationAFBarracks().ArmList...)
	gaouter.as[ukb.Name] = ukb.Architecture
	return ukb.Architecture
}

func aggregationSABarracks() *model.Architecture {
	a := &model.Architecture{
		Id:                sab.Id,
		Name:              sab.Name,
		ArmList:           append(sab.ArmList, sab.barracks.ArmList...),
		BloodVolume:       sab.BloodVolume,
		ConstructionPrice: sab.ConstructionPrice,
	}
	gaouter.as[sab.Name] = a
	return a
}

func aggregationIRAQBarracks() *model.Architecture {
	iraqb.ArmList = append(iraqb.ArmList, aggregationSABarracks().ArmList...)
	gaouter.as[iraqb.Name] = iraqb.Architecture
	return iraqb.Architecture
}

func arlist() {
	aronce.Do(func() {
		var data []model.Architecture
		for _, a := range gaouter.as {
			data = append(data, *a)
		}
		gaouter.list = data
	})
}

func defaultAouter() *aouter {
	ao := &aouter{list: make([]model.Architecture, 0), as: map[string]*model.Architecture{}}
	return ao
}

func (ao *Architecture) ArchitectureList() []model.Architecture {
	return gaouter.list
}

func (ao *Architecture) FetchArchitecture(Name string) (model.Architecture, error) {
	if a, ok := gaouter.as[Name]; ok {
		return *a, nil
	}
	return emptyArchitecture, errors.New("this Architecture Name is not build")
}
