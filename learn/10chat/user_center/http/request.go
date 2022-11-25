package http

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Gctx  *gin.Context
	Error error
	Data  interface{}
	Msg   string
	Code  int
}

func NewRequest(ctx *gin.Context) *Request {
	return &Request{
		Gctx:  ctx,
		Error: errors.New(""),
	}
}

func (r *Request) SetError(err error) *Request {
	r.Error = err
	return r
}

func (r *Request) SetData(data interface{}) *Request {
	r.Data = data
	return r
}

func (r *Request) SetMsg(msg string) *Request {
	r.Msg = msg
	return r
}

func (r *Request) SetbusinessCode(code int) *Request {
	r.Code = code
	return r
}

func (r *Request) SetCtx(ctx *gin.Context) *Request {
	r.Gctx = ctx
	return r
}

func (r *Request) CommonResponse(code int) {
	r.Gctx.JSON(code, gin.H{
		"code":  r.Code,
		"msg":   r.Msg,
		"data":  r.Data,
		"error": r.Error.Error(),
	})
}

func (r *Request) SuccessResponse() {
	r.Gctx.JSON(0, gin.H{
		"code":  200,
		"data":  r.Data,
		"msg":   r.Msg,
		"error": r.Error.Error(),
	})
}
