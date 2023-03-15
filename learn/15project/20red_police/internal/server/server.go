package server

import (
	"20red_police/internal/service"
)

type Server struct {
	UserSrc   *service.UserService
	RoomSrc   *service.RoomService
	PlayerSrc *service.PlayerService
	PMapSrc   *service.PMapService
}

func NewServer(userSrc *service.UserService, roomSrc *service.RoomService, playerSrc *service.PlayerService, pmapSrc *service.PMapService) *Server {
	return &Server{
		UserSrc:   userSrc,
		RoomSrc:   roomSrc,
		PlayerSrc: playerSrc,
		PMapSrc:   pmapSrc,
	}
}
