package protocol

import "20red_police/internal/model"

type Room struct {
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	MapName      string             `json:"map_name"`
	MapUserCount int                `json:"map_user_count"`
	Status       int                `json:"status"`
	CreateTime   string             `json:"create_time"`
	Players      map[string]*Player `json:"players"`
	Owner        string             `json:"owner"`
}

type JoinRoomRequest struct {
	Base
	RoomID string `json:"room_id" mapstructure:"room_id" validate:"required"`
}

type JoinRoomResponse struct {
	Room
}

type CreateRoomRequest struct {
	Base
	RoomName string `json:"room_name" mapstructure:"room_name" validate:"required"`
	PMapID   string `json:"pmap_id" mapstructure:"room_name" validate:"required"`
}

type CreateRoomResponse struct {
	Room
}

type RoomListRequest struct {
	Base
}
type RoomListResponse struct {
	List []Room `json:"list"`
}

func FormatRoomByDBToPro(model *model.Room) Room {
	players := make(map[string]*Player, len(model.Players))
	for _, player := range model.Players {
		item := FormatPlayerByDBToPro(player)
		players[player.Name] = item
	}
	return Room{
		Id:           model.Id,
		Name:         model.Name,
		MapName:      model.MapName,
		MapUserCount: model.MapUserCount,
		Status:       int(model.Status),
		CreateTime:   model.CreateTime,
		Players:      players,
		Owner:        model.Owner,
	}
}
