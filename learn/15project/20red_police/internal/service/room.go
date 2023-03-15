package service

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"fmt"
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

	user, err := rs.userRepo.UserCanTransformPlayer(username)
	if err != nil {
		return nil, err
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
	user.SetPlaying()
	if _, err = rs.userRepo.Update(user); err != nil {
		return nil, err
	}

	fmt.Println("user:update->", user)
	return &nRoom, err
}

func (rs *RoomService) JoinRoom(username string, roomID string) (*model.Room, error) {

	user, err := rs.userRepo.UserCanTransformPlayer(username)
	if err != nil {
		return nil, err
	}

	player := model.NewPlayer(username)
	nRoom, err := rs.roomRepo.JoinRoom(player, roomID)
	if err != nil {
		return nil, err
	}

	user.SetPlaying()
	if _, err = rs.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &nRoom, err
}

func (rs *RoomService) OutRoom(username string, roomID string) (model.Room, error) {

	room, err := rs.roomRepo.FetchRoom(roomID)
	if err != nil {
		return room, err
	}
	if err = rs.roomRepo.OutRoom(username, roomID); err != nil {
		return room, err
	}

	user, err := rs.userRepo.FetchUser(username)
	if err != nil {
		return room, err
	}
	user.SetPrepare()
	if _, err = rs.userRepo.Update(user); err != nil {
		return room, err
	}

	return room, nil
}

func (rs *RoomService) DeleteRoom(username, roomID string) error {
	room, err := rs.roomRepo.FetchRoom(roomID)
	if err != nil {
		return err
	}
	if err = rs.roomRepo.DeleteRoom(username, roomID); err != nil {
		return err
	}

	for playerName, _ := range room.Players {
		user, err := rs.userRepo.FetchUser(playerName)
		if err != nil {
			fmt.Println("delete room , fetchuser err:", err)
			continue
		}
		user.SetPrepare()
		if _, err = rs.userRepo.Update(user); err != nil {
			return err
		}
	}
	return nil
}

func (rs *RoomService) RoomList() ([]model.Room, error) {
	return rs.roomRepo.RoomList(), nil
}
