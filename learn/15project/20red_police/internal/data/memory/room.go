package memory

import (
	"20red_police/internal/model"
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

func init() {
	gonce.Do(func() {
		grooms = &rooms{
			list: map[string]*model.Room{},
		}
	})
}

func (r *Room) Create(room model.Room, username string) model.Room {
	return empty
}

func (r *Room) Dissolve(roomID string, username string) error {
	return nil
}

func (r *Room) Update(room *model.Room) (model.Room, error) {
	return empty, nil
}

func (r *Room) List() []model.Room {
	return nil
}

func (r *Room) JoinRoom(player *model.Player, roomID string) error {
	return nil
}

func (r *Room) OutRoom(player *model.Player, roomID string) error {
	return nil
}

func (r *Room) Broadcast(rootID string) error {
	return nil
}
