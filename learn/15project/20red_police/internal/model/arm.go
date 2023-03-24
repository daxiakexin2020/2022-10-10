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
	grizzly_tank            = "灰熊坦克"
	rhino_tank              = "犀牛坦克"
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

var gengineer *arm = &arm{id: "1", name: engineer, damage_value: damage_value_0, blood_volume: blood_volume_100}

func ProductionEngineer() *arm {
	return gengineer
}

//
//func mjforce() *arm {
//	return &arm{id: "2", name: "盟军大兵", damage_value: 10, blood_volume: 100}
//}
