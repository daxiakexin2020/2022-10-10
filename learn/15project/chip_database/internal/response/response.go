package response

import (
	cerr "chip_database/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponse(cerr.SUCCESS_CODE, "ok", data))
}

func Err500(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, NewResponse(cerr.SERVER_ERR_CODE, "server inner error", nil))
}

func Err404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, NewResponse(cerr.SUCCESS_CODE, "not found", nil))
}

func ParamErr(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponse(cerr.PARAM_ERR_CODE, msg, data))
}

func DBErr(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponse(cerr.DB_ERR_CODE, msg, data))
}

func Err(ctx *gin.Context, code int32, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponse(code, msg, data))
}

func NewResponse(code int32, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
