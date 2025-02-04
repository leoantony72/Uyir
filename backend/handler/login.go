package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/leoantony72/Uyir/model"
	"golang.org/x/crypto/bcrypt"
)

// Session expiry time
const sessionTTL = 24 * time.Hour

func LoginUser(c *gin.Context) {
	loginRequest := model.User{}

	// Parse login request
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid request", "err": err.Error()})
		return
	}

	// Retrieve user data from Redis
	key := "user:" + loginRequest.Name
	ctx := context.Background()
	userData, err := Db.Get(ctx, key).Result()
	if err == redis.Nil {
		c.JSON(401, gin.H{"message": "User not found"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"message": "Error retrieving user data", "err": err.Error()})
		return
	}

	// Unmarshal user data
	var storedUser model.User
	err = json.Unmarshal([]byte(userData), &storedUser)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error unmarshaling user data", "err": err.Error()})
		return
	}

	// Compare stored hashed password with entered password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(401, gin.H{"message": "Invalid credentials"})
		return
	}

	// Generate a session token
	sessionToken := uuid.New().String()

	// Store session in Redis with expiry
	sessionKey := "session:" + sessionToken
	err = Db.Set(ctx, sessionKey, storedUser.Id, sessionTTL).Err()
	if err != nil {
		c.JSON(500, gin.H{"message": "Error creating session", "err": err.Error()})
		return
	}

	// Set session token as a cookie
	// Set session token and user name as cookies
	c.SetCookie("session_token", sessionToken, int(sessionTTL.Seconds()), "/", "", false, true)
	c.SetCookie("user_name", storedUser.Name, int(sessionTTL.Seconds()), "/", "", false, true)

	// Successful login response
	c.JSON(200, gin.H{"message": "Login successful", "user": storedUser.Name})
}
