package controller

import (
	"12lucky_draw/server"
	"github.com/gin-gonic/gin"
)

func ShowTimeDrawPollAction(ctx *gin.Context) {
	ds := server.NewDrawService()
	res := ds.ShowTimeDrawPoll()
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": res,
	})
}

func CountResultAction(ctx *gin.Context) {
	ds := server.NewDrawService()
	res := ds.CountResult()
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": res,
	})
}
