//go:build wireinject
// +build wireinject

package main

import (
	"chip_database/conf"
	"chip_database/internal/data/db"
	"chip_database/internal/data/http"
	kernal "chip_database/internal/kernel"
	"chip_database/internal/route"
	"chip_database/internal/server"
	"chip_database/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initApp(eneing *gin.Engine, c *conf.WebServerConfig, sqliteC *conf.SqliteDatabaseConfig, fileHandleProxyC *conf.FileHandleProxyConfig) (*kernal.Kernel, error) {
	panic(wire.Build(route.ProviderSet, server.ProviderSet, service.ProviderSet, kernal.ProviderSet, db.ProviderSet, http.ProviderSet))
}
