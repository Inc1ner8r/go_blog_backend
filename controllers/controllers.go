package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/models"
)

func CreateBlog(c *gin.Context) {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	var blog models.Blog

	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
}
