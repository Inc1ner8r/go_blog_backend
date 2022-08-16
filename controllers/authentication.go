package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/inciner8r/blog_backend_go/models"
)

var jwt_key = []byte("secret_key")

var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var credentials Credentials
	var expected Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	if err := db.Table("users").Where("username = ?", credentials.Username).First(&expected).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	if expected.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	expirationTime := time.Now().Add(time.Hour * 2)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
	}

	c.SetCookie("token", tokenString, int(expirationTime.Hour()-time.Now().Hour()), "/", "localhost", false, true)
}

// func Login(c *gin.Context) {
// 	var credentials Credentials
// 	if err := c.BindJSON(&credentials); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
// 		return
// 	}

// 	expectedPassword, ok := Users[credentials.Username]

// 	if !ok || expectedPassword != credentials.Password {
// 		c.JSON(http.StatusUnauthorized, gin.H{"data": "unauthorized"})
// 	}

// 	expirationTime := time.Now().Add(time.Hour * 2)

// 	claims := &Claims{
// 		Username: credentials.Username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)t
// 	tokenString, err := token.SignedString(jwt_key)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
// 	}
// 	cookie, err := c.Cookie("token")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
// 	}

// 	c.SetCookie(cookie, tokenString, int(expirationTime.Hour()-time.Now().Hour()), "/", "localhost", false, true)
// }
