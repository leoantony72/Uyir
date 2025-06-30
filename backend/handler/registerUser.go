package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password", "error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	user.Verified = false // Always set to false on registration

	if user.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Role is required"})
		return
	}

	switch user.Role {
	case "user":
		if user.FirstName == "" || user.LastName == "" || user.Name == "" || user.Email == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing required fields for user"})
			return
		}
		user.Address = ""
		user.Phone = ""
		user.Verified = true
	case "hospital", "police", "pwd":
		if user.Name == "" || user.Email == "" || user.Password == "" || user.Address == "" || user.Phone == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing required fields for " + user.Role})
			return
		}
		// User-specific fields should be empty/zero
		user.Points = 0
		user.VechicleType = ""
		user.FuelType = ""
		user.VechicleNumber = ""
		user.FirstName = ""
		user.LastName = ""
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid role"})
		return
	}

	result := Db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving user", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered", "id": user.ID, "role": user.Role})
}
