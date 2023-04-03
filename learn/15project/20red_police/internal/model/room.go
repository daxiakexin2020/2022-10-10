package model

import (
	"20red_police/tools"
	"errors"
	"fmt"
	"sync"
)

type status int

const (
	STATUS_WAITING status = iota + 1
	STATUS_PLAYING
	STATUS_OVER
	STATUC_FULL
)

var STATUS_MAPING = map[status]string{
	STATUS_WAITING: "等待中",
	STATUS_PLAYING: "已开始",
	STATUS_OVER:    "已结束",
	STATUC_FULL:    "已满员",
}

type Room struct {
	Id           string
	Name         string
	MapName      string
	MapUserCount int
	Status       status
	CreateTime   string
	Players      map[string]*Player
	Owner        string
	InitPirce    int32
	mu           sync.RWMutex `json:"-"`
}

func NewRoom(roomName, username, pmapName string, count int, initPrice int32) *Room {
	return &Room{
		//Id:           tools.UUID(),
		Id:           "1",
		Name:         roomName,
		MapName:      pmapName,
		MapUserCount: count,
		Status:       STATUS_WAITING,
		CreateTime:   tools.NowTimeFormatTimeToString(),
		Players:      map[string]*Player{},
		Owner:        username,
		InitPirce:    initPrice,
	}
}

func (r *Room) JoinRoom(player *Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.isCanJoin(); err != nil {
		return err
	}
	if _, ok := r.Players[player.Name]; ok {
		return errors.New("you already join this room,can not repeat join")
	}
	r.Players[player.Name] = player
	if len(r.Players) == r.MapUserCount {
		r.Status = STATUC_FULL
	}
	return nil
}

func (r *Room) OutRoom(playerName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.isCanOutRoom(); err != nil {
		return err
	}
	if _, ok := r.Players[playerName]; !ok {
		return errors.New("you are out this room")
	}
	delete(r.Players, playerName)
	if r.Owner == playerName {
		r.Status = STATUS_OVER
	} else {
		r.Status = STATUS_WAITING
	}
	return nil
}

func (r *Room) DeleteRoom(playerName string) error {
	r.mu.Lock()
	r.mu.Unlock()
	if err := r.isCanOutRoom(); err != nil {
		return err
	}
	if r.Owner != playerName {
		return errors.New("room ower can start delete room")
	}
	r.Status = STATUS_OVER
	return nil
}

func (r *Room) RoomStatus() status {
	return r.Status
}

func (r *Room) GameStart(username string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	players := r.Players
	if len(players) <= 1 {
		return errors.New("this room has one player，geme can not start")
	}
	if r.Owner != username {
		return errors.New("room ower can start game only")
	}
	for _, player := range players {
		if !player.IsReady() {
			return errors.New("has player no ready")
		}
	}
	r.Status = STATUS_PLAYING
	return nil
}

func (r *Room) UpdateRoomPlayer(playerName string, status bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.Status == STATUS_PLAYING {
		return errors.New("game is playing ,can not update player status")
	}
	player, ok := r.Players[playerName]
	if !ok {
		return fmt.Errorf("this player:%s not in this room:%s", playerName, r.Id)
	}
	player.Status = status
	return nil
}

func (r *Room) IsOver() bool {
	return r.Status == STATUS_OVER
}

func (r *Room) isCanJoin() error {
	if r.Status == STATUS_WAITING {
		return nil
	}
	return errors.New("join failed，game" + STATUS_MAPING[r.Status])
}

func (r *Room) isCanOutRoom() error {
	if r.Status == STATUS_PLAYING {
		return errors.New("game" + STATUS_MAPING[r.Status])
	}
	return nil
}

func (r *Room) IsCanCreate() bool {
	return r.Id != "" && r.Name != ""
}

func (r *Room) UserIsInRoom(username string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if _, ok := r.Players[username]; ok {
		return true
	}
	return false
}

func (r *Room) FetchPlayer(username string) (*Player, error) {
	if p, ok := r.Players[username]; ok {
		return p, nil
	}
	return nil, errors.New("this room has no this user")
}
