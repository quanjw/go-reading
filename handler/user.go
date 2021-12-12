package handler

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.DefaultPostForm("password", "123456")
	passwordAgain := c.DefaultPostForm("password-again", "123456")
	println("email", email, "password", password, "password again", passwordAgain)
	c.JSON(200, gin.H{
		"message": "用户" + username + "已注册",
		"success": "true",
	})
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
