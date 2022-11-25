package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"user_center/http"
)

func Test(ctx *gin.Context) {
	r := http.NewRequest(ctx)
	r.SetData("test").SetbusinessCode(10001).SetError(errors.New("test error")).CommonResponse(300)
}

func Test2(ctx *gin.Context) {
	r := http.NewRequest(ctx)
	r.SuccessResponse()
}
