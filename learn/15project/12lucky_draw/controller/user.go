package controller

import (
	"12lucky_draw/server"
	"github.com/gin-gonic/gin"
)

func AddAction(ctx *gin.Context) {
	username := ctx.PostForm("username")
	us := server.NewUserService()
	user, err := us.Add(username)
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"err":  errStr,
		"data": user,
	})
}

func UpdateAction(ctx *gin.Context) {}

func DeleteAction(ctx *gin.Context) {}

func FetchAction(ctx *gin.Context) {
	uid := ctx.Query("uid")
	us := server.NewUserService()
	user, err := us.Find(uid)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"err":  err,
		"data": user,
	})
}

func FtechAllAction(ctx *gin.Context) {
	us := server.NewUserService()
	data := us.All()
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func DrawAction(ctx *gin.Context) {
	uid := ctx.PostForm("uid")
	us := server.NewUserService()
	level, err := us.Draw(uid)
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"err":  errStr,
		"data": level,
	})
}
