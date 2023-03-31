package model

import (
	"20red_police/tools"
	"sync"
	"sync/atomic"
)

type Player struct {
	Id                  string
	Name                string
	CountryName         string
	Color               string
	Money               int32
	Mch                 chan int32
	Status              bool
	OutCome             bool
	Architectures       map[string]*alItem
	Arms                map[string]*armItem
	IsBuildingAR        bool
	buildingARMu        sync.RWMutex `json:"-"`
	buildingARMuThredId string
}

type alItem struct {
	name  string
	count uint32
}

type armItem struct {
	name  string
	count uint32
}

func NewPlayer(name string) *Player {
	p := &Player{
		Id:            tools.UUID(),
		Name:          name,
		Architectures: map[string]*alItem{},
		Mch:           make(chan int32, 1000),
	}
	p.Mch <- 10000 //init money
	return p
}

func (p *Player) hadnleMoney(money int32) {
	atomic.AddInt32(&p.Money, money)
}

func (p *Player) HandleMoney(money int32) {
	p.hadnleMoney(money)
}

func (p *Player) PickFormCh() int32 {
	m := <-p.Mch
	return m
}

func (p *Player) AddToCh(money int32) {
	p.Mch <- money
}

func (p *Player) closeMch() {
	close(p.Mch)
}

func (p *Player) CloseMch() {

}

func (p *Player) IsReady() bool {
	return p.Status
}
func (p *Player) SetReady() {
	p.Status = true
}

func (p *Player) Country() string {
	return p.CountryName
}

func (p *Player) PMoney() int32 {
	return p.Money
}

func (p *Player) SetUnReady() {
	p.Status = false
}

func (p *Player) SetIsBuildingAR(threadId string) bool {
	p.buildingARMu.Lock()
	defer p.buildingARMu.Unlock()
	if p.IsBuildingAR {
		return false
	}
	p.IsBuildingAR = true
	p.buildingARMuThredId = threadId
	return true
}

func (p *Player) SetIsNotBuildingAR(threaId string) bool {
	if threaId != p.buildingARMuThredId {
		return false
	}
	p.buildingARMu.RLock()
	defer p.buildingARMu.RUnlock()
	if !p.IsBuildingAR {
		return false
	}
	p.IsBuildingAR = false
	return true
}
