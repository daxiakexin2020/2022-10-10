package controller

import (
	cvalidator "14gateway/components/validator"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Name       string `json:"name" form:"name" uri:"name" validate:"required,len=10"`
	Password   string `json:"password" form:"password" uri:"password"`
	Code       string `json:"code" form:"code" uri:"code" validate:"required"`
	IsDisabled bool   `json:"is_disabled" form:"is_disabled" uri:"is_disabled"`
}

func Tes(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": "0",
		"msg":  "ok",
	})
}

// body
func TestJsonHandler(ctx *gin.Context) {

	login := &Login{}
	err := ctx.ShouldBindJSON(login)
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

// form
func TestFormHandler(ctx *gin.Context) {
	login := &Login{}
	err := ctx.Bind(login)
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
