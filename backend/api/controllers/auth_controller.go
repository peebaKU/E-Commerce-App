package controllers

import (
	"ecommerce-app/database"
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register a new user
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login user and return JWT token
func Login(c *gin.Context) {
	var user models.User
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Where("email = ?", input.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
