package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
	}
	return db
}

var db = initDB()

func CreateBlog(c *gin.Context) {

	var blog models.Blog
	results, err := db.Query("select * from test_table")

	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
}
