package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "root:root@tcp(db:3306)/test?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Blog{})
	fmt.Println("db init")
	return db
}

var db = initDB()

func CreateBlog(c *gin.Context) {
	var blog models.Blog

	if err := c.BindJSON(&blog); err != nil {
		fmt.Println("err1")
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	if err := db.Create(&blog).Error; err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": blog})
}

func DisplayBlogs(c *gin.Context) {
	var blogs []models.Blog

	if err := db.Find(&blogs).Error; err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, blogs)
}

func GetBlog(c *gin.Context) {
	var blog models.Blog

	if err := db.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": blog})
}
