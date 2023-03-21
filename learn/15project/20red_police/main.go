package main

import (
	"20red_police/asynchronous"
	"20red_police/asynchronous/room_timeout"
	"20red_police/config"
	iasynchronous "20red_police/internal/asynchronous"
	"20red_police/internal/middleware"
	"20red_police/internal/synchronization/file/stores"
	"20red_police/network"
	"log"
	"os"
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

	go handleExit()

	roomTimeout := room_timeout.Timeout(10000, time.Second*time.Duration(config.GetRoomConfig().RoomLiveTime), iasynchronous.Handle)
	if err := asynchronous.GoAsynchronous(roomTimeout); err != nil {
		panic(err)
	}

	if err := stores.GoSynchronizationRunBuilder(); err != nil {
		panic(err)
	}

	network.RegisterMiddleware(middleware.LogMiddleware, middleware.LoginGuardMiddleware, middleware.ValidatorMiddleWare)
	network.Run()
}

func handleExit() {

	defer func() {
		os.Exit(0)
	}()

	<-network.GOEXIT

	if err := asynchronous.STOP(); err != nil {
		log.Println("stop asynchronous err, please handle:", err)
	}

	if err := stores.GoSynchronizationStopBuilder(); err != nil {
		log.Println("GoSynchronizationStopBuilder err:", err)
	}
}
