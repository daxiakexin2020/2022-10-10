package router

import (
	"f_gin/controller"
	"f_gin/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(route *gin.Engine) {
	route.Use(middleware.TokenMiddleware(), middleware.LogMiddleware())
	route.POST("/test_json", controller.TestJsonHandler)
	route.POST("/test_form", controller.TestFormHandler)
	route.POST("/test_uri/:name/:code/:password", controller.TestUriHandler)
	route.GET("/test_query", controller.TestQueryHandler)
}
