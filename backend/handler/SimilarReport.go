package handler

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/corona10/goimagehash"
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/model"
	"path/filepath"

	"image"
	_ "image/jpeg"
	_ "image/png"
)

// Define input structure
type Input struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Haversine formula to calculate distance in meters
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Earth radius in meters
	lat1, lon1, lat2, lon2 = lat1*(math.Pi/180), lon1*(math.Pi/180), lat2*(math.Pi/180), lon2*(math.Pi/180)
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func computeHash(filePath string) (*goimagehash.ImageHash, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return goimagehash.PerceptionHash(img)
}

// SimilarReports function
func SimilarReports(c *gin.Context) {
	// Parse form values
	latStr := c.PostForm("latitude")
	lonStr := c.PostForm("longitude")
	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid latitude"})
		return
	}
	longitude, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid longitude"})
		return
	}

	// Handle image file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image file is required"})
		return
	}

	// Save to temporary path
	tempPath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save uploaded image"})
		return
	}
	defer os.Remove(tempPath) // clean up temp file

	// Compute perceptual hash for input image
	inputHash, err := computeHash(tempPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to process input image"})
		return
	}

	// Fetch pending reports
	var reports []model.Report
	if err := Db.Table("reports").Where("status = ?", "Pending").Find(&reports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching reports"})
		return
	}

	// Compare with reports
	var similarReports []model.Report
	for _, report := range reports {
		//500m proximity
		distance := haversine(latitude, longitude, report.Latitude, report.Longitude)
		if distance > 500 {
			continue
		}

		if report.FilePath == "" {
			continue
		}

		hash, err := computeHash(report.FilePath)
		if err != nil {
			fmt.Printf("Error hashing image for report %s: %v\n", report.ID, err)
			continue
		}

		//hash distance (<=10 similar)
		difference, err := inputHash.Distance(hash)
		if err != nil {
			continue
		}
		if difference <= 10 {
			similarReports = append(similarReports, report)
		}
	}
	c.JSON(http.StatusOK, gin.H{"similar_reports": similarReports})
}