package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-reading/conf"
	"log"
	"net/http"
	"strings"
)

func LoginAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		/*cookie, err := context.Request.Cookie("gin_user")
		if err == nil {
			//name, value string, maxAge int, path, domain string, secure, httpOnly bool
			context.SetCookie(cookie.Name, cookie.Value, 7*24*3600, "/", conf.Domain, false, true)
			context.Next()
		} else {
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}*/

		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "无法认证，重新登录",
				"success": "false",
			})
			context.Abort()
			return
		}
		auth = strings.Fields(auth)[1]

		claims, err := parseToken(auth)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "无法认证，重新登录",
				"success": "false",
			})
			context.Abort()
			return
		} else {
			log.Println("token正常", claims)
		}

	}
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.SECRET), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
