package middleware

import (
	"github.com/gin-gonic/gin"
	"go-reading/conf"
	"net/http"
)

func LoginAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Request.Cookie("gin_user")
		if err == nil {
			//name, value string, maxAge int, path, domain string, secure, httpOnly bool
			context.SetCookie(cookie.Name, cookie.Value, 7*24*3600, "/", conf.Domain, false, true)
			context.Next()
		} else {
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}
