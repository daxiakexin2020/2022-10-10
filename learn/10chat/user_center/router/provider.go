package router

import (
	"github.com/gin-gonic/gin"
	"user_center/http"
)

func BootRouter() error {
	r := gin.Default()
	InitRouter(r)
	http.R = r
	return nil
}
