package memory

import "20red_police/internal/model"

type Player struct{}

func (p *Player) Create(name string) model.Player {
	return model.Player{}
}
