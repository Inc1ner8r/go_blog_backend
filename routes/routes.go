package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "api is up",
		})
	})

	router.POST("/newBlog", controllers.CreateBlog)
	router.GET("/getBlogs", controllers.DisplayBlogs)
	router.GET("/getBlog", controllers.getBlog)
}
