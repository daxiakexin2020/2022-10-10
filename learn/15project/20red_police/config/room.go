package config

import (
	"log"
	"sync"
)

var (
	roomConfig RoomConfig
	ronce      sync.Once
)

type RoomConfig struct {
	RoomLiveTime int64 `json:"room_live_time"`
}

func (rc *RoomConfig) CName() string {
	return "room"
}

func makeRoomConfig() RoomConfig {
	ronce.Do(func() {
		rc := RoomConfig{}
		if err := Generate(rc.CName(), &rc); err != nil {
			log.Fatalf("读取%s配置出错%v", rc.CName(), err)
		}
		roomConfig = rc
	})
	return roomConfig
}

func GetRoomConfig() RoomConfig {
	return makeRoomConfig()
}
