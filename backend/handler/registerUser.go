package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/model"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registers a new user and saves it to PostgreSQL
func RegisterUser(c *gin.Context) {
	var user model.User

	// Bind JSON input to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password", "error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)

	// Save user to PostgreSQL
	result := Db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving user", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered", "id": user.ID})
}
