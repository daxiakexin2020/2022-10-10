package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"errors"
	"sync"
)

type Room struct{}

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
		return *room, errors.New("create room faield")
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

func (r *Room) RoomList() []model.Room {
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
		return emptyRoom, errors.New("no this room:" + roomID)
	}
	if err := room.JoinRoom(player); err != nil {
		return emptyRoom, err
	}
	return *room, nil
}

func (r *Room) FetchRoom(roomID string) (model.Room, error) {
	grooms.mu.RLock()
	defer grooms.mu.RUnlock()
	if room, ok := grooms.list[roomID]; ok {
		return *room, nil
	}
	return emptyRoom, errors.New("no this room:" + roomID)
}

func (r *Room) OutRoom(playerName string, roomID string) error {
	grooms.mu.Lock()
	defer grooms.mu.Unlock()
	room, ok := grooms.list[roomID]
	if !ok {
		return errors.New("no this room:" + roomID)
	}
	if err := room.OutRoom(playerName); err != nil {
		return err
	}
	return nil
}

func (r *Room) DeleteRoom(playerName string, roomID string) error {
	grooms.mu.Lock()
	defer grooms.mu.Unlock()
	room, ok := grooms.list[roomID]
	if !ok {
		return errors.New("no this room:" + roomID)
	}
	if err := room.DeleteRoom(playerName); err != nil {
		return err
	}
	delete(grooms.list, roomID)
	return nil
}

func (r *Room) GameStart(username, roomID string) error {
	grooms.mu.Lock()
	defer grooms.mu.Unlock()
	room, ok := grooms.list[roomID]
	if !ok {
		return errors.New("no this room:" + roomID)
	}
	if err := room.GameStart(username); err != nil {
		return err
	}
	return nil
}

func (r *Room) UpdateRoomPlayer(roomID string, username string, status bool) error {
	grooms.mu.Lock()
	defer grooms.mu.Unlock()
	room, ok := grooms.list[roomID]
	if !ok {
		return errors.New("no this room:" + roomID)
	}
	return room.UpdateRoomPlayer(username, status)
}

func (r *Room) Broadcast(rootID string) error {
	return nil
}
