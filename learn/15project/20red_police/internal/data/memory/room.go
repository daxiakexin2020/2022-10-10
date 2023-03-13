package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"errors"
	"sync"
)

type Room struct {
}

type rooms struct {
	list map[string]*model.Room
	mu   sync.RWMutex
}

var (
	ronce  sync.Once
	empty  model.Room
	grooms *rooms
)

var _ data.Room = (*Room)(nil)

var emptyRoom = model.Room{}

func init() {
	gonce.Do(func() {
		grooms = &rooms{
			list: map[string]*model.Room{},
		}
	})
}

func NewRoom() data.Room {
	return &Room{}
}

func (r *Room) Create(room *model.Room) (model.Room, error) {
	grooms.mu.Lock()
	defer grooms.mu.Unlock()
	if !room.IsCanCreate() {
		return *room, errors.New("创建房间失败")
	}
	grooms.list[room.Id] = room
	return *room, nil
}

func (r *Room) Dissolve(roomID string, username string) error {
	return nil
}

func (r *Room) Update(room *model.Room) (model.Room, error) {
	return empty, nil
}

func (r *Room) List() []model.Room {
	var res []model.Room
	for _, room := range grooms.list {
		res = append(res, *room)
	}
	return res
}

func (r *Room) JoinRoom(player *model.Player, roomID string) (model.Room, error) {
	grooms.mu.Lock()
	defer grooms.mu.Unlock()
	room, ok := grooms.list[roomID]
	if !ok {
		return emptyRoom, errors.New("房间不存在")
	}
	if err := room.JoinRoom(player); err != nil {
		return emptyRoom, err
	}
	return *room, nil
}

func (r *Room) OutRoom(player *model.Player, roomID string) error {
	return nil
}

func (r *Room) Broadcast(rootID string) error {
	return nil
}
