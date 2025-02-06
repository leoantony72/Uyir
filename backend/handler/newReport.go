package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

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

	// Parse form data to handle file upload
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB limit
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	// Handle file upload
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}
	defer file.Close()

	// Generate a unique filename
	filename := uuid.New().String() + filepath.Ext(header.Filename)
	filepath := "uploads/" + filename

	// Save the file
	if err := c.SaveUploadedFile(header, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Bind JSON form data to report
	report.UserID = session.UserID
	report.ID = uuid.New().String()
	report.FilePath = filepath

	fmt.Println("Report:", report)

	// Save report to the database
	if err := Db.Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report submitted", "user_id": report.UserID, "file_path": report.FilePath})
}
