package model

import (
	"20red_police/components/operation"
	"20red_police/tools"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Player struct {
	Id            string
	Name          string
	CountryName   string
	Color         string
	Money         int32
	Mch           chan int32
	Status        bool
	OutCome       bool
	Architectures map[string]*arItem
	Arms          map[string]*armItem

	isBuildingAR        bool
	buildingARMuThredId string
	addARMu             sync.RWMutex `json:"-"`
	buildingARMu        sync.RWMutex `json:"-"`

	isBuildingARM        bool
	buildingARMMuThredId string
	buildingARMMu        sync.RWMutex `json:"-"`
	addARMMu             sync.RWMutex `json:"-"`
}

type arItem struct {
	name  string
	count uint32
}

type armItem struct {
	name  string
	count uint32
}

const minInitPrice = 100

var EXIT_PLAYER = make(chan struct{}, 1)

func init() {
	operation.Oeration().Register(EXIT_PLAYER)
}

func NewPlayer(name string, initPrice int32) *Player {
	p := &Player{
		Id:            tools.UUID(),
		Name:          name,
		Architectures: map[string]*arItem{},
		Arms:          map[string]*armItem{},
		Mch:           make(chan int32, 1000),
	}
	if initPrice < minInitPrice {
		initPrice = minInitPrice
	}
	p.Mch <- initPrice
	p.Money = initPrice
	return p
}

func (p *Player) addArchitecture(name string) {
	p.addARMu.RLock()
	defer p.addARMu.RUnlock()
	item, ok := p.Architectures[name]
	if !ok {
		nitem := p.newArItem(name)
		p.Architectures[name] = nitem
	} else {
		item.count++
	}
}

func (p *Player) addArm(name string) {
	p.addARMMu.RLock()
	defer p.addARMMu.RUnlock()
	item, ok := p.Arms[name]
	if !ok {
		nitem := p.newArmItem(name)
		p.Arms[name] = nitem
	} else {
		item.count++
	}
}

func (p *Player) newArItem(name string) *arItem {
	return &arItem{name: name, count: 1}
}

func (p *Player) newArmItem(name string) *armItem {
	return &armItem{name: name, count: 1}
}

func (p *Player) PickFormCh() int32 {
	m := <-p.Mch
	atomic.AddInt32(&p.Money, -m)
	return m
}

func (p *Player) addToCh(money int32) {
	p.Mch <- money
	atomic.AddInt32(&p.Money, money)
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

func (p *Player) SetBuildingAR(threadId string) bool {
	p.buildingARMu.Lock()
	defer p.buildingARMu.Unlock()
	if p.isBuildingAR {
		return false
	}
	p.isBuildingAR = true
	p.buildingARMuThredId = threadId
	return true
}

func (p *Player) setNotBuildingAR(threaId string) bool {
	if threaId != p.buildingARMuThredId {
		return false
	}
	p.buildingARMu.RLock()
	defer p.buildingARMu.RUnlock()
	if !p.isBuildingAR {
		return false
	}
	p.isBuildingAR = false
	return true
}

func (p *Player) SetBuildingARM(threadId string) bool {
	p.buildingARMMu.Lock()
	defer p.buildingARMMu.Unlock()
	if p.isBuildingARM {
		return false
	}
	p.isBuildingARM = true
	p.buildingARMMuThredId = threadId
	return true
}

func (p *Player) setNotBuildingARM(threaId string) bool {
	if threaId != p.buildingARMMuThredId {
		return false
	}
	p.buildingARMMu.RLock()
	defer p.buildingARMMu.RUnlock()
	if !p.isBuildingARM {
		return false
	}
	p.isBuildingARM = false
	return true
}

func (p *Player) BuildAROver(price int32, threadId string, arName string) {
	p.addToCh(price)
	p.setNotBuildingAR(threadId)
	p.addArchitecture(arName)
}

func (p *Player) BuildARMOver(price int32, threadId string, armName string) {
	p.addToCh(price)
	p.setNotBuildingARM(threadId)
	p.addArm(armName)
}

func (p *Player) MockAddPirceIntoCh() {
	timer := time.NewTicker(time.Second * 60)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			p.addToCh(1000)
		default:
			if len(EXIT_PLAYER) == 1 {
				fmt.Println("我们退出了")
			}
		}
	}
}
