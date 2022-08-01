package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/controllers"
)

func routes(router *gin.Engine) {
	router.POST("/newBlog", controllers.CreateBlog)
}
