package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CostTimeLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		nowTime := time.Now()
		context.Next()
		costTime := time.Since(nowTime)
		url := context.Request.URL.String()
		fmt.Printf("the request url %s cost %v \n", url, costTime)
	}
}
