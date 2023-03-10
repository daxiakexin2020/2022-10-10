package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
)

type Player struct{}

var _ data.Player = (*Player)(nil)

func NewPlayer() data.Player {
	return &Player{}
}

func (p *Player) Create(name string) model.Player {
	return model.Player{}
}
