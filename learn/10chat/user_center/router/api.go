package router

import (
	"github.com/gin-gonic/gin"
	"user_center/http/controller"
)

func InitRouter(r *gin.Engine) {
	r.GET("/test", controller.Test)
	r.GET("/test2", controller.Test2)
}
