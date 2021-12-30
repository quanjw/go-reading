package router

import (
	"github.com/gin-gonic/gin"
	"go-reading/handler"
	"go-reading/middleware"
	"go-reading/utils"
	"net/http"
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

	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CostTimeLog())

	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/go-reading/templates/*"))
	r.Static("/statics", "./statics")

	r.GET("/ping", ping)

	r.GET("/index", middleware.LoginAuth(), handler.Index)

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.Login)
	}

	noteRouter := r.Group("/note")
	{
		noteRouter.POST("/upload", handler.UploadNote)
		noteRouter.POST("/insert", handler.NoteInsert)
	}

	r.StaticFS("/upload", http.Dir(utils.RootPath()+"upload/"))

	adminRouter := r.Group("adminm")
	{
		adminRouter.Use(gin.BasicAuth(gin.Accounts{
			"admin": "123456",
		}))
		adminRouter.GET("index", handler.AdminIndex)
	}

	return r
}
