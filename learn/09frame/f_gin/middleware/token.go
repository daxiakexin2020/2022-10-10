package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("***************token middleware****************")

		//中断
		//context.Abort()

		//继续
		context.Next()
	}
}
