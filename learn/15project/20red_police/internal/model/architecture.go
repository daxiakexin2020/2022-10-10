package model

import "sync"

type Architecture struct {
	Id                string
	Name              string
	ArmList           []string
	BloodVolume       int
	ConstructionPrice int32
	CanBeBuildNum     int
	CanSell           bool
	CanPutAway        bool
	mu                sync.RWMutex `json:"-"`
}

func (a *Architecture) FetchArchitectureBloodVolume() int {
	return a.BloodVolume
}

func (a *Architecture) FetchArchitectureArm() []string {
	return a.ArmList
}

func (a *Architecture) FetchArchitectureConstructionPrice() int32 {
	return a.ConstructionPrice
}
