package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/routes"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	routes.Test()
}
