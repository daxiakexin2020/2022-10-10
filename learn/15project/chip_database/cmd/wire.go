//go:build wireinject
// +build wireinject

package main

import (
	"chip_database/conf"
	kernal "chip_database/internal/kernel"
	"chip_database/internal/route"
	"chip_database/internal/server"
	"chip_database/internal/service"
	"github.com/gin-gonic/gin"
)

func initApp(eneing *gin.Engine, c *conf.WebServerConfig) (*kernal.Kernel, error) {
	panic(wire.Build(route.ProviderSet, server.ProviderSet, service.ProviderSet, kernal.ProviderSet))
}
