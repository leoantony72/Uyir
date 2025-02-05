package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leoantony72/Uyir/model"
)

func NewReport(c *gin.Context) {
	var report model.Report
	cookie, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing session token"})
		return
	}

	fmt.Println("Session Token:", cookie)

	var session model.Session
	if err := Db.Where("token = ?", cookie).First(&session).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
		return
	}

	fmt.Println("User ID:", session.UserID)
	if err := c.BindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	report.UserID = session.UserID
	report.ID = uuid.New().String()
	fmt.Println("Report:", report)
	
	if err := Db.Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report submitted", "user_id": report.UserID})
}
