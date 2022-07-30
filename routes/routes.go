package routes

import "github.com/gin-gonic/gin"

func routes(router *gin.Engine) {
	router.POST("/newBlog", controllers.CreateBlog)
}
