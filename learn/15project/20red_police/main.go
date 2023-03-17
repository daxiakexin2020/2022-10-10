package main

import (
	"20red_police/asynchronous"
	"20red_police/asynchronous/room_timeout"
	"20red_police/config"
	iasynchronous "20red_police/internal/asynchronous"
	"20red_police/internal/middleware"
	"20red_police/network"
	"time"
)

func main() {
	run()
}

func run() {

	s := initApp()
	if err := network.Register(s); err != nil {
		panic(err)
	}
	if err := config.InitializeProxyViper(); err != nil {
		panic(err)
	}
	roomTimeout := room_timeout.Timeout(10000, time.Second*time.Duration(config.GetRoomConfig().RoomLiveTime), iasynchronous.Handle)
	if err := asynchronous.GoAsynchronous(roomTimeout); err != nil {
		panic(err)
	}

	network.RegisterMiddleware(middleware.LogMiddleware, middleware.LoginGuardMiddleware, middleware.ValidatorMiddleWare)
	network.Run()
}
