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
	user.Role = "user" // Default role for new users
	user.Points = 0 // Initialize points to 0
	
	result := Db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving user", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered", "id": user.ID})
}


func RegisterHospital(c *gin.Context) {
    var hospital model.Hospitals
    if err := c.ShouldBindJSON(&hospital); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
        return
    }
    hospital.Verified = false 
    result := Db.Create(&hospital)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving hospital", "error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Hospital registered", "id": hospital.ID})
}

func RegisterPolice(c *gin.Context) {
    var police model.PoliceStation
    if err := c.ShouldBindJSON(&police); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
        return
    }
    police.Verified = false // Default to not verified
    result := Db.Create(&police)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving police station", "error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Police station registered", "id": police.ID})
}

func RegisterPWD(c *gin.Context) {
    var pwd model.PWD
    if err := c.ShouldBindJSON(&pwd); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
        return
    }
    pwd.Verified = false // Default to not verified
    result := Db.Create(&pwd)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving PWD", "error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "PWD registered", "id": pwd.ID})
}