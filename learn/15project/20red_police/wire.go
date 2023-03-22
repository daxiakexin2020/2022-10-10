//go:build wireinject
// +build wireinject

package main

import (
	"20red_police/internal/data/memory"
	"20red_police/internal/server"
	"20red_police/internal/service"
	"20red_police/network"
	"github.com/google/wire"
)

// wire.go
func initApp() *server.Server {
	panic(wire.Build(memory.ProviderSet, service.ProviderSet, network.ProviderSet, server.ProviderSet))
}
