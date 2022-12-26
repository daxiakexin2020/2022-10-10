package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
)

/**
https://blog.csdn.net/asd1126163471/article/details/123453278
*/

func LimiterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := rate.NewLimiter(0, 100)
		pctx, _ := context.WithTimeout(context.Background(), time.Second*1)
		err := limiter.Wait(pctx)
		if err != nil {
			fmt.Println("被限流")
			ctx.Abort()
		}
		fmt.Println("limiter", limiter)
		ctx.Next()
	}
}
