package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		defer func() {
			s := time.Now().Sub(startTime).String()
			log.Printf("COST TIME:::::::%s\n", s)
		}()
		remoteAddr := context.Request.RemoteAddr
		url := context.Request.URL
		body := context.Request.GetBody
		log.Printf("request info, remoteAddr：%s，url：%s，body：%v\n", remoteAddr, url, body)
		context.Next()
	}
}
