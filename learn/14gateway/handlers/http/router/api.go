package router

import (
	"14gateway/handlers/http/controller"
	"14gateway/handlers/http/middleware"
	"github.com/gin-gonic/gin"
)

var E *gin.Engine

func InitApiRouter() error {
	E = gin.New()
	E.Use(middleware.ParseMiddleWare())
	api := E.Group("/")
	api.Any("/test", controller.Tes)
	api.POST("/test_json", controller.TestJsonHandler)
	api.POST("/test_form", controller.TestFormHandler)
	api.POST("/test_uri/:name/:code/:password", controller.TestUriHandler)
	api.GET("/test_query", controller.Tes)
	return nil
}
