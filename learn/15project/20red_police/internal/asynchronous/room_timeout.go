package asynchronous

import (
	"20red_police/internal/data"
	"20red_police/internal/data/memory"
	"20red_police/internal/model"
	"errors"
	"fmt"
	"log"
)

func HandleRoomTimeout(roomID string) error {
	room, err := deleteRoom(roomID)
	if err != nil {
		return err
	}
	if err = updateUserStatus(room); err != nil {
		return err
	}
	return nil
}

func deleteRoom(roomID string) (*model.Room, error) {
	pick, err := data.MemoryRoom()
	if err != nil {
		return nil, err
	}
	room, ok := pick.(*memory.Room)
	if !ok {
		return nil, errors.New("memory class is err")
	}
	fetchRoom, err := room.FetchRoom(roomID)
	if err != nil {
		log.Println("fetch room err:", err)
		return nil, err
	}
	if err = fetchRoom.DeleteRoom(fetchRoom.Owner); err != nil {
		log.Println("delete room err:", err)
		return nil, err
	}
	fmt.Println("delete room ok:", roomID)
	return &fetchRoom, nil
}

func updateUserStatus(room *model.Room) error {
	pick, err := data.MemoryUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*memory.User)
	if !ok {
		return errors.New("memory class is err")
	}
	for playerName, _ := range room.Players {
		muser, err := user.FetchUser(playerName)
		if err != nil {
			fmt.Println("update User status , fetchuser err:", err)
			continue
		}
		muser.SetPrepare()
		if _, err = user.Update(muser); err != nil {
			return err
		}
	}
	log.Println("update room player status ok")
	return nil
}
