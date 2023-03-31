package server

import (
	"20red_police/internal/service"
	"20red_police/network"
)

type Server struct {
	UserSrc    *service.UserService
	RoomSrc    *service.RoomService
	PlayerSrc  *service.PlayerService
	PMapSrc    *service.PMapService
	NetworkSrc *network.Resources
	CountrySrc *service.CountryService
}

func NewServer(userSrc *service.UserService, roomSrc *service.RoomService, playerSrc *service.PlayerService, pmapSrc *service.PMapService, networkSrc *network.Resources, countrySrc *service.CountryService) *Server {
	return &Server{
		UserSrc:    userSrc,
		RoomSrc:    roomSrc,
		PlayerSrc:  playerSrc,
		PMapSrc:    pmapSrc,
		NetworkSrc: networkSrc,
		CountrySrc: countrySrc,
	}
}
