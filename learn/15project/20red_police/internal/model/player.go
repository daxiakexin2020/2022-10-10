package model

import (
	"20red_police/tools"
	"log"
	"sync"
	"time"
)

type Player struct {
	Id                   string
	Name                 string
	CountryName          string
	Color                string
	Money                int32
	Status               bool
	OutCome              bool
	Architectures        map[string]*arItem
	Arms                 map[string]*armItem
	mch                  chan int32
	gameOverCh           chan struct{}
	isBuildingAR         bool
	buildingARMuThredId  string
	addARMu              sync.RWMutex
	buildingARMu         sync.RWMutex
	isBuildingARM        bool
	buildingARMMuThredId string
	buildingARMMu        sync.RWMutex
	addARMMu             sync.RWMutex
	gameOverMu           sync.RWMutex
}

type arItem struct {
	name  string
	count uint32
}

type armItem struct {
	name  string
	count uint32
}

const (
	minInitPrice     = 100
	BuildARTimeout   = 20 * 60
	BuildARMTimeout  = 20 * 60
	MockAddPirceTime = 20
	mchCap           = 10000
)

func NewPlayer(name string, initPrice int32) *Player {
	p := &Player{
		Id:            tools.UUID(),
		Name:          name,
		Architectures: map[string]*arItem{},
		Arms:          map[string]*armItem{},
		mch:           make(chan int32, mchCap),
		gameOverCh:    make(chan struct{}, 1),
	}
	if initPrice < minInitPrice {
		initPrice = minInitPrice
	}
	p.mch <- initPrice
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

func (p *Player) addToCh(money int32) {
	p.mch <- money
}

func (p *Player) closeMch() {
	close(p.mch)
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
	p.buildingARMu.Lock()
	defer p.buildingARMu.Unlock()
	if threaId != p.buildingARMuThredId {
		return false
	}
	if !p.isBuildingAR {
		return false
	}
	p.isBuildingAR = false
	p.buildingARMuThredId = ""
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
	p.buildingARMMu.Lock()
	defer p.buildingARMMu.Unlock()
	if threaId != p.buildingARMMuThredId {
		return false
	}
	if !p.isBuildingARM {
		return false
	}
	p.isBuildingARM = false
	p.buildingARMMuThredId = ""
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
	timer := time.NewTicker(time.Second * MockAddPirceTime)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			p.addToCh(1000)
		default:
			if len(p.gameOverCh) == 1 {
				log.Println("game over,MockAddPirceIntoCh out.................")
				return
			}
		}
	}
}

func (p *Player) gameOver() {
	p.gameOverCh <- struct{}{}
}

func (p *Player) GameOver() {
	p.gameOverMu.RLock()
	defer p.gameOverMu.RUnlock()
	if len(p.gameOverCh) == 1 {
		return
	}
	p.gameOver()
}

func (p *Player) LenGameOverch() int {
	return len(p.gameOverCh)
}

func (p *Player) PickReadMch() <-chan int32 {
	return p.mch
}

func (p *Player) PickWriteMch() chan<- int32 {
	return p.mch
}
