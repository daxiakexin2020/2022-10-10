package service

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"errors"
)

type RoomService struct {
	roomRepo   data.Room
	userRepo   data.User
	playerRepo data.Player
}

func NewRoomService(roomRepo data.Room, userRepo data.User, playerRepo data.Player) *RoomService {
	return &RoomService{roomRepo: roomRepo, userRepo: userRepo, playerRepo: playerRepo}
}

func (rs *RoomService) CreateRoom(roomName, username, pmapID string, count int) (*model.Room, error) {
	if _, ok := rs.userRepo.IsOnLine(username); !ok {
		return nil, errors.New("用户离线")
	}
	room, err := rs.roomRepo.Create(model.NewRoom(roomName, username, pmapID, count))
	if err != nil {
		return nil, err
	}

	player := model.NewPlayer(username)
	nRoom, err := rs.roomRepo.JoinRoom(player, room.Id)
	if err != nil {
		return nil, err
	}

	return &nRoom, err
}

func (rs *RoomService) JoinRoom(user model.User, roomID string) (*model.Room, error) {
	player := model.NewPlayer(user.Name)
	nRoom, err := rs.roomRepo.JoinRoom(player, roomID)
	if err != nil {
		return nil, err
	}
	return &nRoom, err
}

func (rs *RoomService) RoomList() ([]model.Room, error) {
	return rs.roomRepo.List(), nil
}
