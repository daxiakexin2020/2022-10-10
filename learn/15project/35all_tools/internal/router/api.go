package router

import (
	"35all_tools/internal/handlers"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	engine      *gin.Engine
	jsonHandler *handlers.JsonHanler
}

func NewApiRouter(engine *gin.Engine, jh *handlers.JsonHanler) *ApiRouter {
	return &ApiRouter{
		engine:      engine,
		jsonHandler: jh,
	}
}

func (ar *ApiRouter) RegisterHandlers() {
	jsonGroup := ar.engine.Group("/json/")
	{
		jsonGroup.POST("check_json", ar.jsonHandler.CheckJson)
		jsonGroup.POST("json_to_golang_struct", ar.jsonHandler.JsonToGolangStruct)
	}
}
