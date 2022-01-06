package main

import (
	"go-reading/router"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name quanjw
// @contact.url https://wwww.example.com
// @contact.email admin@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	r := router.SetupRouter()
	_ = r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
