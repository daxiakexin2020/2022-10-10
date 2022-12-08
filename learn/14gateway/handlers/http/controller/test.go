package controller

import (
	cvalidator "14gateway/components/validator"
	"14gateway/handlers/http/service"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Name       string `json:"name" form:"name" uri:"name" validate:"required,len=10"`
	Password   string `json:"password" form:"password" uri:"password"`
	Code       string `json:"code" form:"code" uri:"code" validate:"required"`
	IsDisabled bool   `json:"is_disabled" form:"is_disabled" uri:"is_disabled"`
}

func Tes(ctx *gin.Context) {

	service := service.NewClientServer(ctx)
	resp, err := service.Do()

	if err != nil {
		ctx.JSON(200, gin.H{
			"code": "1002",
			"msg":  "请求失败",
			"data": ctx.Request,
			"err":  err.Error(),
		})
	} else {
		ctx.JSON(200, resp)
	}
}

// body
func TestJsonHandler(ctx *gin.Context) {

	service := service.NewClientServer(ctx)
	resp, err := service.Do()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":  "1002",
			"msg":   "请求失败",
			"data":  ctx.Request.URL,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, resp)
	}
}

// form
func TestFormHandler(ctx *gin.Context) {
	service := service.NewClientServer(ctx)
	resp, err := service.Do()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":  "1002",
			"msg":   "请求失败",
			"data":  ctx.Request.URL,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, resp)
	}
}

// param
func TestUriHandler(ctx *gin.Context) {
	login := &Login{}
	err := ctx.ShouldBindUri(login)
	if err != nil {
		ctx.JSON(0, gin.H{
			"code": 1001,
			"msg":  err.Error(),
			"data": login,
		})
		return
	}

	verr := cvalidator.Check(login)
	if verr != nil {
		ctx.JSON(0, gin.H{
			"code": 1002,
			"msg":  verr.Error(),
			"data": login,
		})
		return
	}

	ctx.JSON(0, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": login,
	})
}

// query
func TestQueryHandler(ctx *gin.Context) {

	login := &Login{}
	err := ctx.ShouldBindQuery(login)
	if err != nil {
		ctx.JSON(0, gin.H{
			"code": 1001,
			"msg":  err.Error(),
			"data": login,
		})
		return
	}

	verr := cvalidator.Check(login)
	if verr != nil {
		ctx.JSON(0, gin.H{
			"code": 1002,
			"msg":  verr.Error(),
			"data": login,
		})
		return
	}

	ctx.JSON(0, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": login,
	})
}
