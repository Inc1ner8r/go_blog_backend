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

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Blog{})
	fmt.Println("db init")
	return db
}

var db = InitDB()

func CreateBlog(c *gin.Context) {
	var blog models.Blog

	if err := c.BindJSON(&blog); err != nil {
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

func UpdateBlog(c *gin.Context) {
	var blog models.Blog
	var inputBlog models.Blog

	if err := db.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//validate input
	if err := c.ShouldBindJSON(&inputBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	db.Model(&blog).Updates(inputBlog)
	c.JSON(http.StatusOK, gin.H{"data": blog})
}

func DeleteBlog(c *gin.Context) {
	var blog models.Blog

	if err := db.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	db.Delete(&blog)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
