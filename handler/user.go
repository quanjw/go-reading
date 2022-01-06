package handler

import (
	"github.com/gin-gonic/gin"
	"go-reading/conf"
	"go-reading/model"
	"log"
	"net/http"
	"strconv"
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
			c.JSON(200, gin.H{
				"message": "用户" + user.Username + "已登陆",
				"success": "true",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "登陆失败，用户名或密码错误",
				"success": "true",
			})
		}

	}

}
