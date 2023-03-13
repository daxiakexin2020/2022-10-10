package service

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
)

type PlayerService struct {
	data data.Player
}

func NewPlayerService(data data.Player) *PlayerService {
	return &PlayerService{
		data: data,
	}
}

func (ps *PlayerService) Create(name string) (*model.Player, error) {
	player := model.NewPlayer(name)
	return player, nil
}
