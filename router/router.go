package router

import (
	"github.com/gin-gonic/gin"
	"go-reading/handler"
	"os"
	"path/filepath"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"success": "true",
	})
}

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/go-reading/templates/*"))
	r.Static("/statics", "./statics")

	r.GET("/ping", ping)

	r.GET("/index", handler.Index)

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.Login)
	}

	noteRouter := r.Group("/note")
	{
		noteRouter.POST("/upload", handler.UploadNote)
	}

	return r
}
