package router

import (
	"12lucky_draw/controller"
	"github.com/gin-gonic/gin"
)

func InitApi(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/add", controller.AddAction)
		user.POST("/update", controller.UpdateAction)
		user.DELETE("/delete", controller.DeleteAction)
		user.GET("/fetch", controller.FetchAction)
		user.GET("fetchall", controller.FtechAllAction)
		user.POST("/draw", controller.DrawAction)
	}
	draw := router.Group("/draw")
	{
		draw.GET("show_time_draw_poll", controller.ShowTimeDrawPollAction)
		draw.GET("count_result", controller.CountResultAction)
	}
}
