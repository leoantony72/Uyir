package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leoantony72/Uyir/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Session expiry time
const sessionTTL = 24 * time.Hour

// LoginUser logs a user in and sets a session
func LoginUser(c *gin.Context) {
	loginRequest := model.User{}

	// Parse login request
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "err": err.Error()})
		return
	}

	// Retrieve user data from PostgreSQL
	var storedUser model.User
	if err := Db.Where("name = ?", loginRequest.Name).First(&storedUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user data", "err": err.Error()})
		}
		return
	}

	// Compare stored hashed password with entered password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Generate a session token
	sessionToken := uuid.New().String()

	// Store session token in the database (assuming you have a sessions table in PostgreSQL)
	session := model.Session{
		Token:     sessionToken,
		UserID:    storedUser.ID,
		ExpiresAt: time.Now().Add(sessionTTL),
	}

	if err := Db.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating session", "err": err.Error()})
		return
	}

	// Set session token and user name as cookies
	c.SetCookie("session_token", sessionToken, int(sessionTTL.Seconds()), "/", "", false, true)
	c.SetCookie("user_name", storedUser.Name, int(sessionTTL.Seconds()), "/", "", false, true)

	// Successful login response
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": storedUser.Name})
}
