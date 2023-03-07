package service

import "20red_police/internal/data"

type PlayerService struct {
	data data.Player
}

func NewPlayer(data data.Player) *PlayerService {
	return &PlayerService{
		data: data,
	}
}
