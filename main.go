package main

import (
	"go-reading/router"
)

func main() {
	r := router.SetupRouter()
	_ = r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
