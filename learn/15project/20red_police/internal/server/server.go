package server

import (
	"20red_police/config"
	"20red_police/internal/service"
	"20red_police/tools"
	"errors"
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

func (s *Server) check(token string, dest string) error {
	consumerData, err := tools.ParseToken(token, config.GetJwtConfig().TokenSecret)
	if err != nil {
		return err
	}
	v, ok := consumerData.(*interface{})
	if !ok {
		return errors.New("consumer data is err in token ")
	}
	tname := *v
	if _, ok = tname.(string); !ok {
		return errors.New("consumer data type is error")
	}
	if tname != dest {
		return errors.New("consumer data is not eq dest!!!!!! ")
	}
	return nil
}
