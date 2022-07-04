package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func init() {
	conn, err := GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	DB = conn
}

func main() {
	r := gin.Default()

	r.GET("/", index)
	r.POST("/register", register)
	r.POST("/login", login)
	r.GET("/private", Authorized(), private)

	err := r.Run(":9000")
	if err != nil {
		log.Fatalln(err)
	}
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "index route",
	})
}

func register(c *gin.Context) {
	var req Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var newUser User
	// Hash password
	hashedPass, _ := newUser.HashPassword(req.Password)

	newUser = User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPass,
	}

	// Insert to db
	err := DB.Debug().Create(&newUser).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "user successfully registered",
	})
}

func login(c *gin.Context) {
	var req Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var user User
	err := DB.Debug().Where("email", req.Email).Find(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})

		return
	}

	pass := user.CheckPassword(req.Password, user.Password)
	if pass != true {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "username or password not correct",
		})

		return
	}

	token, _ := user.GenerateJWT(req.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func private(c *gin.Context) {
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "private route",
		"user": email,
	})
}
