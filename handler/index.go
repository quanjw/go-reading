package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "go-reading",
	})
}
