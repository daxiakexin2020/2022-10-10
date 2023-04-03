package model

import "sync"

type Arm struct {
	Id                string
	Name              string
	DamageValue       int
	BloodVolume       int
	ConstructionPrice int32
	mu                sync.RWMutex `json:"-"`
}

func (a *Arm) FetchArchitectureBloodVolume() int {
	return a.BloodVolume
}

func (a *Arm) FetchArmConstructionPrice() int32 {
	return a.ConstructionPrice
}
