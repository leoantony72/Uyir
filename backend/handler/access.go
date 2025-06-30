package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/model"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil || sessionToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Session token missing"})
			c.Abort()
			return
		}

		// Find session in DB
		var session model.Session
		if err := Db.Where("token = ?", sessionToken).First(&session).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session"})
			c.Abort()
			return
		}
		if session.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Session expired"})
			c.Abort()
			return
		}

		var user model.User
		if err := Db.Where("id = ?", session.UserID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			c.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			if strings.EqualFold(user.Role, allowed) {
				c.Set("user", user)
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden: insufficient permissions"})
		c.Abort()
	}
}
