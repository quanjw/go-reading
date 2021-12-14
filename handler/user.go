package handler

import (
	"github.com/gin-gonic/gin"
	"go-reading/model"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		c.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		println("username", user.Username, "email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		c.JSON(200, gin.H{
			"message": "用户" + user.Username + "已注册",
			"success": "true",
		})
		//c.Redirect(http.StatusMovedPermanently, "/")重定向
	}

}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	//password := c.PostForm("password")
	//email := c.PostForm("email")
	//username := c.PostForm("username")
	//username := c.PostForm("username")
	c.JSON(200, gin.H{
		"message": "用户" + username + "已登陆",
		"success": "true",
	})
}
