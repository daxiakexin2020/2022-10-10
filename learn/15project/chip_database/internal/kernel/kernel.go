package kernal

import (
	"chip_database/conf"
	"chip_database/internal/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Kernel struct {
	serverConfig *conf.WebServerConfig
	engine       *gin.Engine
	route        *route.ApiRouter
}

func New(engine *gin.Engine, serverConfig *conf.WebServerConfig, route *route.ApiRouter) *Kernel {
	return &Kernel{
		engine:       engine,
		serverConfig: serverConfig,
		route:        route,
	}
}

func (k *Kernel) Run() {
	k.route.RegisterHandlers()
	addr := fmt.Sprintf("%s:%d", k.serverConfig.Host, k.serverConfig.Port)
	if err := k.engine.Run(addr); err != nil {
		log.Fatalf("server run err:%v\n", err)
	}
}

func (k *Kernel) Stop() {}
