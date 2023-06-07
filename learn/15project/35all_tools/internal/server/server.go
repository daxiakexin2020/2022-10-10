package server

import (
	"35all_tools/conf"
	"35all_tools/internal/model"
	"35all_tools/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

type server struct {
	name      string
	host      string
	port      int
	engine    *gin.Engine
	apiRouter *router.ApiRouter
}

var _ (model.ServerRepo) = (*server)(nil)

func NewServer(engine *gin.Engine, apiRouter *router.ApiRouter, conf *conf.WebServerConfig) model.ServerRepo {
	return &server{
		name:      conf.Name,
		host:      conf.Host,
		port:      conf.Port,
		engine:    engine,
		apiRouter: apiRouter,
	}
}

func (s *server) DefaultServerRun() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	return s.run(addr)
}

func (s *server) Run(addr string) error {
	return s.run(addr)
}

func (s *server) Stop() error {
	return nil
}

func (s *server) run(addr string) error {
	s.apiRouter.RegisterHandlers()
	return s.engine.Run(addr)
}
