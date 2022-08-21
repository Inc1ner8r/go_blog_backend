package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/inciner8r/blog_backend_go/models"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return
	}

	user.Password = string(hash)

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

	if err := bcrypt.CompareHashAndPassword([]byte(expected.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
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

	c.SetCookie("jwt", tokenString, int(expirationTime.Hour()-time.Now().Hour()), "/", "localhost", false, true)
	fmt.Println(c.Cookie("jwt"))
}

func ValidateJWT(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "cookie not found"})
		return
	}
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}
	claims := token.Claims.(*Claims)
	fmt.Println(claims.Username)
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
