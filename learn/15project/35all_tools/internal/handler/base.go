package handler

import (
	cerr "35all_tools/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Base struct{}

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewBase() *Base {
	return &Base{}
}

func (b *Base) ShouldBindJSON(ctx *gin.Context, dest interface{}) error {
	return ctx.ShouldBindJSON(&dest)
}

func (b *Base) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponse(cerr.SUCCESS_CODE, "ok", data))
}

func (b *Base) Err500(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, NewResponse(cerr.SERVER_ERR_CODE, "server inner error", nil))
}

func (b *Base) Err404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, NewResponse(cerr.SUCCESS_CODE, "not found", nil))
}

func (b *Base) Err(ctx *gin.Context, code int32, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponse(code, msg, data))
}

func NewResponse(code int32, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
