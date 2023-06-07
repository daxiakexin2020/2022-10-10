//go:build wireinject
// +build wireinject

package main

import (
	"35all_tools/conf"
	"35all_tools/internal/data/local"
	"35all_tools/internal/handlers"
	"35all_tools/internal/model"
	"35all_tools/internal/router"
	"35all_tools/internal/server"
	"35all_tools/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp(engine *gin.Engine, conf *conf.WebServerConfig) (model.ServerRepo, error) {
	panic(wire.Build(server.ProviderSet, router.ProviderSet, handlers.ProviderSet, service.ProviderSet, local.ProviderSet))
}
