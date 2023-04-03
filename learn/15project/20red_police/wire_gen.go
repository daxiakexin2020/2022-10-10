// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"20red_police/internal/data/memory"
	"20red_police/internal/server"
	"20red_police/internal/service"
	"20red_police/network"
)

// Injectors from wire.go:

// wire.go
func initApp() *server.Server {
	user := memory.NewUser()
	userService := service.NewUserService(user)
	room := memory.NewRoom()
	player := memory.NewPlayer()
	roomService := service.NewRoomService(room, user, player)
	architecture := memory.NewArchitecture()
	arm := memory.NewArm()
	playerService := service.NewPlayerService(player, room, architecture, arm)
	pMap := memory.NewPMap()
	pMapService := service.NewPMapService(pMap)
	resources := network.Gresources()
	country := memory.NewCountry()
	countryService := service.NewCountryService(country)
	serverServer := server.NewServer(userService, roomService, playerService, pMapService, resources, countryService)
	return serverServer
}
