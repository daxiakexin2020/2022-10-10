package model

import (
	"errors"
	"sync"
)

type architecture struct {
	id                string
	name              string
	armList           []string
	bloodVolume       int
	constructionPrice int
}

type aouter struct {
	list          []string
	architectures map[string]*architecture
}

type barracks struct {
	*architecture
}

type afbarracks struct {
	*architecture
	*barracks
}

type ukBarracks struct {
	*architecture
	*afbarracks
}

type saBarracks struct {
	*architecture
	*barracks
}

type iraqBarracks struct {
	*architecture
	*saBarracks
}

var b = &barracks{
	&architecture{
		id:      "1",
		name:    barrack,
		armList: []string{engineer},
	},
}

var afb = &afbarracks{
	architecture: &architecture{
		id:      "2",
		name:    af_barrack,
		armList: []string{af_soldier, af_flying_soldier, af_spy, af_dog, af_chronosphere_soldier},
	},
	barracks: b,
}

var ukb = &ukBarracks{
	architecture: &architecture{
		id:      "3",
		name:    uk_barrack,
		armList: []string{british_sniper},
	},
	afbarracks: afb,
}

var sab = &saBarracks{
	architecture: &architecture{
		id:      "4",
		name:    sa_barrack,
		armList: []string{sa_soldier, sa_dog},
	},
	barracks: b,
}

var iraqb = &iraqBarracks{
	architecture: &architecture{
		id:      "5",
		name:    iraq_barrack,
		armList: []string{iraq_badiation_engineer},
	},
	saBarracks: sab,
}

const (
	barrack      = "兵营"
	af_barrack   = "盟军兵营"
	uk_barrack   = "英国兵营"
	sa_barrack   = "苏军兵营"
	iraq_barrack = "伊拉克兵营"
)

var (
	ionce   sync.Once
	gaouter *aouter
)

func InitArchitecture() *aouter {
	ionce.Do(func() {
		gaouter = defaultAouter()
		aggregationAFBarracks()
		aggregationUKBarracks()
		aggregationSABarracks()
		aggregationIRAQBarracks()
		list()
	})
	return gaouter
}

func Gaouter() *aouter {
	if gaouter == nil {
		return InitArchitecture()
	}
	return gaouter
}

/**
  build
*/

func aggregationAFBarracks() *architecture {
	a := &architecture{
		id:          afb.id,
		name:        afb.name,
		armList:     append(afb.armList, afb.barracks.armList...),
		bloodVolume: afb.bloodVolume,
	}
	gaouter.architectures[afb.name] = a
	return a
}

func aggregationUKBarracks() *architecture {
	ukb.armList = append(ukb.armList, aggregationAFBarracks().armList...)
	gaouter.architectures[ukb.name] = ukb.architecture
	return ukb.architecture
}

func aggregationSABarracks() *architecture {
	a := &architecture{
		id:          sab.id,
		name:        sab.name,
		armList:     append(sab.armList, sab.barracks.armList...),
		bloodVolume: sab.bloodVolume,
	}
	gaouter.architectures[sab.name] = a
	return a
}

func aggregationIRAQBarracks() *architecture {
	iraqb.armList = append(iraqb.armList, aggregationSABarracks().armList...)
	gaouter.architectures[iraqb.name] = iraqb.architecture
	return iraqb.architecture
}

func list() {
	var data []string
	for name, _ := range gaouter.architectures {
		data = append(data, name)
	}
	gaouter.list = data
}

func defaultAouter() *aouter {
	ao := &aouter{list: make([]string, 0), architectures: map[string]*architecture{}}
	return ao
}

func (ao *aouter) List() []string {
	return ao.list
}

func (ao *aouter) FetchArchitectureArm(name string) ([]string, error) {
	if a, ok := ao.architectures[name]; ok {
		return a.armList, nil
	}
	return []string{}, errors.New("this architecture name is not build")
}

func (ao *aouter) FetchArchitectureBloodVolume(name string) (int, error) {
	if a, ok := ao.architectures[name]; ok {
		return a.bloodVolume, nil
	}
	return 0, errors.New("this architecture name is not build")
}
