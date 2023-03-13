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

func (p *Player) Create(mode model.Player) model.Player {
	return model.Player{}
}
