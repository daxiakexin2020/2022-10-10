package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"errors"
	"sync"
)

type Arm struct{}

type armOuter struct {
	list []model.Arm
	arms map[string]*model.Arm
}

const (
	engineer                = "工程师"
	af_soldier              = "盟军大兵"
	af_dog                  = "盟军狗"
	af_flying_soldier       = "盟军飞行兵"
	af_chronosphere_soldier = "盟军超时空兵"
	af_spy                  = "盟军间谍"
	af_mine_car             = "mjkc"
	sa_soldier              = "苏军大兵"
	sa_dog                  = "苏军狗"
	sa_mine_car             = "苏军矿车"
	british_sniper          = "英国狙击手"
	iraq_badiation_engineer = "伊拉克辐射工兵"
	grizzly_tank            = "灰熊坦克"
	rhino_tank              = "犀牛坦克"
	tank_slayer             = "坦克杀手"
	light_edge_tank         = "光棱坦克"
	apocalypse_tank         = "天启坦克"
	kirov                   = "基洛夫"
)

const (
	DamageValue_0   = 0
	DamageValue_100 = 100
	DamageValue_200 = 200
	DamageValue_300 = 300
)

const (
	BloodVolume_100 = 100
	BloodVolume_200 = 200
	BloodVolume_300 = 300
	BloodVolume_400 = 400
	BloodVolume_500 = 500
	BloodVolume_600 = 600
)

var emptyArm model.Arm

var (
	gengineer    = &model.Arm{Id: "1", Name: engineer, DamageValue: DamageValue_0, BloodVolume: BloodVolume_100}
	gaf_soldier  = &model.Arm{Id: "2", Name: af_soldier, DamageValue: DamageValue_100, BloodVolume: BloodVolume_100}
	gaf_dog      = &model.Arm{Id: "3", Name: af_dog, DamageValue: DamageValue_0, BloodVolume: BloodVolume_100}
	gaf_mine_car = &model.Arm{Id: "4", Name: af_mine_car, DamageValue: DamageValue_100, BloodVolume: BloodVolume_600, ConstructionPrice: C_ConstructionPrice_500}
)

var (
	_         data.Arm = (*Arm)(nil)
	garmOuter *armOuter
	armOnce   sync.Once
)

func init() {
	garmOuter = defaultArmOuter()
	aggregationGEngineer()
	aggregationGAFSoldier()
	aggregationGAFDog()
	aggregationGAFMineCar()
	armlist()
}

func armlist() {
	armOnce.Do(func() {
		var data []model.Arm
		for _, a := range garmOuter.arms {
			data = append(data, *a)
		}
		garmOuter.list = data
	})
}

func aggregationGEngineer() {
	garmOuter.arms[gengineer.Name] = gengineer
}

func aggregationGAFSoldier() {
	garmOuter.arms[gengineer.Name] = gengineer
}

func aggregationGAFDog() {
	garmOuter.arms[gengineer.Name] = gengineer
}

func aggregationGAFMineCar() {
	garmOuter.arms[gaf_mine_car.Name] = gaf_mine_car
}

func defaultArmOuter() *armOuter {
	armo := &armOuter{list: make([]model.Arm, 0), arms: map[string]*model.Arm{}}
	return armo
}

func NewArm() data.Arm {
	return &Arm{}
}

func (a *Arm) IsMineCar(name string) bool {
	if name == af_mine_car || name == sa_mine_car {
		return true
	}
	return false
}

func (a *Arm) ArmList() []model.Arm {
	return garmOuter.list
}

func (a *Arm) FetchArm(Name string) (model.Arm, error) {
	if ar, ok := garmOuter.arms[Name]; ok {
		return *ar, nil
	}
	return emptyArm, errors.New("this Arm Name is not a arm build")
}
