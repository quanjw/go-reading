package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-reading/docs"
	"go-reading/handler"
	"go-reading/handler/note"
	"go-reading/middleware"
	"go-reading/utils"
	"net/http"
	"os"
	"path/filepath"
)

// PingExample godoc
// @Summary ping pong
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"success": "true",
	})
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func SetupRouter() *gin.Engine {

	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CostTimeLog())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
			eg.GET("/ping", ping)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/go-reading/templates/*"))
	r.Static("/statics", "./statics")

	r.GET("/index", middleware.LoginAuth(), handler.Index)

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.Login)
	}

	noteRouter := r.Group("")
	{
		noteRouter.PUT("/note/:id", note.UploadNote)
		noteRouter.GET("/note/:id", note.Get)
		noteRouter.GET("/notes", note.GetAll)
		noteRouter.POST("/note", note.Insert)
		noteRouter.DELETE("/note/:id", note.Delete)
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
