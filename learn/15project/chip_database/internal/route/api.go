package route

import (
	"chip_database/internal/server"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	engine *gin.Engine
	info   *server.Info
}

func NewApiRouter(engine *gin.Engine, info *server.Info) *ApiRouter {
	return &ApiRouter{engine: engine, info: info}
}

func (r *ApiRouter) RegisterHandlers() {

	api := r.engine.Group("/api/")
	{
		api.GET("list", r.info.List)
	}
}
