package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/model"
)

func GetUnverifiedUsers(c *gin.Context) {
	var users []model.User
	if err := Db.Where("verified = ?", false).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching users", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}


func VerifyUser(c *gin.Context) {
    type Request struct {
        ID string `json:"id"`
    }
    var req Request
    if err := c.ShouldBindJSON(&req); err != nil || req.ID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
        return
    }

    var user model.User
    if err := Db.Where("id = ?", req.ID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }

    user.Verified = true
    if err := Db.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User verified", "id": user.ID})
}