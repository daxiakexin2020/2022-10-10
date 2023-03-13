package model

import (
	"20red_police/tools"
	"errors"
	"sync"
)

type status int

const (
	STATUS_WAITING status = iota + 1
	STATUS_PLAYING
	STATUS_DISSOLVE
	STATUS_OVER
	STATUC_FULL
)

var STATUS_MAPING = map[status]string{
	STATUS_WAITING:  "等待中",
	STATUS_PLAYING:  "已开始",
	STATUS_DISSOLVE: "已解散",
	STATUS_OVER:     "已结束",
	STATUC_FULL:     "已满员",
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
	Mu           sync.Mutex `json:"-"`
}

func NewRoom(roomName, username, pmapName string, count int) *Room {
	return &Room{
		Id:           tools.UUID(),
		Name:         roomName,
		MapName:      pmapName,
		MapUserCount: count,
		Status:       STATUS_WAITING,
		CreateTime:   tools.NowTimeFormatTimeToString(),
		Players:      map[string]*Player{},
		Owner:        username,
	}
}

func (r *Room) JoinRoom(player *Player) error {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	if err := r.isCanJoin(); err != nil {
		return err
	}

	if _, ok := r.Players[player.Name]; ok {
		return errors.New("你已经加入此房间，不可重复加入")
	}
	r.Players[player.Name] = player
	if len(r.Players) == r.MapUserCount {
		r.Status = STATUC_FULL
	}

	return nil
}

func (r *Room) RoomStatus() status {
	return r.Status
}

func (r *Room) IsDissolve() bool {
	return r.Status == STATUS_DISSOLVE
}

func (r *Room) IsOver() bool {
	return r.Status == STATUS_OVER
}

func (r *Room) isCanJoin() error {
	if r.Status == STATUS_WAITING {
		return nil
	}
	return errors.New("加入失败，游戏" + STATUS_MAPING[r.Status])
}

func (r *Room) IsCanCreate() bool {
	return r.Id != "" && r.Name != ""
}
