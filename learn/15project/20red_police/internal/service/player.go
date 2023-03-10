package service

import (
	"20red_police/internal/data"
)

type PlayerService struct {
	data data.Player
}

func NewPlayerService(data data.Player) *PlayerService {
	return &PlayerService{
		data: data,
	}
}
