package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/routes"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	routes.Routes(r)
	r.Run()
}
