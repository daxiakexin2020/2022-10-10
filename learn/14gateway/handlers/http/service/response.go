package service

import (
	proxy_http "14gateway/proxy/http"
	"github.com/gin-gonic/gin"
)

func CommonResponse(resp *proxy_http.ProxyResponse, cs *clientServer) {

	err := cs.GoErr.Unwrap()
	if err != nil {
		cs.ctx.JSON(200, gin.H{
			"code": cs.GoErr.Code,
			"msg":  cs.GoErr.Msg,
			"err":  err.Error(),
		})
	} else {
		cs.ctx.JSON(200, resp)
	}
}
