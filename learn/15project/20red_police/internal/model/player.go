package model

import "20red_police/tools"

type Player struct {
	Id          string
	Name        string
	CountryName string
	Color       string
	Status      bool
	OutCome     bool //结局
}

func NewPlayer(name string) *Player {
	return &Player{
		Id:   tools.UUID(),
		Name: name,
	}
}

func (p *Player) IsReady() bool {
	return p.Status
}
func (p *Player) SetReady() {
	p.Status = true
}

func (p *Player) SetUnReady() {
	p.Status = false
}
