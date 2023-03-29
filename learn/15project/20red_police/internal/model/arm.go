package model

type arm struct {
	id           string
	name         string
	damage_value int
	blood_volume int
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
	damage_value_0   = 0
	damage_value_100 = 100
	damage_value_200 = 200
	damage_value_300 = 300
)

const (
	blood_volume_100 = 100
	blood_volume_200 = 200
	blood_volume_300 = 300
	blood_volume_400 = 400
	blood_volume_500 = 500
	blood_volume_600 = 600
)

var (
	gengineer   = &arm{id: "1", name: engineer, damage_value: damage_value_0, blood_volume: blood_volume_100}
	gaf_soldier = &arm{id: "2", name: af_soldier, damage_value: damage_value_100, blood_volume: blood_volume_100}
	gaf_dog     = &arm{id: "3", name: af_dog, damage_value: damage_value_0, blood_volume: blood_volume_100}
)

func ProductionEngineer() *arm {
	return gengineer
}
func ProductionAFSoldier() *arm {
	return gaf_soldier
}
func ProductionAFDog() *arm {
	return gaf_dog
}
