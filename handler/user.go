package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-reading/conf"
	"go-reading/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		c.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		isExit := user.ExistUser()
		var data map[string]interface{}
		if isExit {
			data = gin.H{
				"message": "用户已存在",
				"success": "true",
			}
		} else {
			id := user.Save()
			data = gin.H{
				"message": "用户" + user.Username + "已注册 ID :" + strconv.FormatInt(id, 10),
				"success": "true",
			}
		}
		c.JSON(200, data)
		//c.Redirect(http.StatusMovedPermanently, "/")重定向
	}

}

// Login PingExample godoc
// @Summary user login
// @Schemes
// @Description user login
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} success
// @Router /user/login [post]
func Login(c *gin.Context) {
	var loginUser model.LoginUser
	if err := c.Bind(&loginUser); err != nil {
		log.Println("err ->", err.Error())
		c.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		user := loginUser.QueryByEmail()
		if loginUser.Password == user.Password {
			c.SetCookie("gin_user", strconv.FormatInt(user.Id, 10), 7*24*3600, "/", conf.Domain, false, true)

			expiresTime := time.Now().Unix() + int64(conf.OneDayOfHours)
			calims := jwt.StandardClaims{
				Audience:  user.Username,     // 受众
				ExpiresAt: expiresTime,       // 失效时间
				Id:        string(user.Id),   // 编号
				IssuedAt:  time.Now().Unix(), // 签发时间
				Issuer:    "quanjw",          // 签发人
				NotBefore: time.Now().Unix(), // 生效时间
				Subject:   "login",           // 主题
			}
			tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, calims)
			jwtSecret := []byte(conf.SECRET)
			if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
				c.JSON(200, gin.H{
					"message": "用户" + user.Username + "已登陆",
					"success": "true",
					"data":    "Bearer " + token,
				})
			} else {
				c.JSON(200, gin.H{
					"message": "生成token错误",
					"success": "true",
				})
			}

		} else {
			c.JSON(200, gin.H{
				"message": "登陆失败，用户名或密码错误",
				"success": "true",
			})
		}

	}

}
