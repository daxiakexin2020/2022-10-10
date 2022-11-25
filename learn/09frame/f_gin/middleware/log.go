package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("***************log middleware****************")

		//中断
		//context.Abort()

		//继续
		context.Next()
	}
}
