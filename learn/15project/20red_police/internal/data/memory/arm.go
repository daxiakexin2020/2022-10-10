package memory

import "20red_police/internal/model"

type Arm struct {
}

const (
	engineer                = "工程师"
	af_soldier              = "盟军大兵"
	af_dog                  = "盟军狗"
	af_flying_soldier       = "盟军飞行兵"
	af_chronosphere_soldier = "盟军超时空兵"
	af_spy                  = "盟军间谍"
	sa_soldier              = "苏军大兵"
	sa_dog                  = "苏军狗"
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

var (
	gengineer   = &model.Arm{Id: "1", Name: engineer, DamageValue: DamageValue_0, BloodVolume: BloodVolume_100}
	gaf_soldier = &model.Arm{Id: "2", Name: af_soldier, DamageValue: DamageValue_100, BloodVolume: BloodVolume_100}
	gaf_dog     = &model.Arm{Id: "3", Name: af_dog, DamageValue: DamageValue_0, BloodVolume: BloodVolume_100}
)
