package main

import (
	"time"

	"github.com/leoantony72/Uyir/AI"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/handler"

)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173","http://localhost:8000"}, // Change this if your frontend URL changes
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", handler.GetData)

	r.POST("/signup", handler.RegisterUser)

	r.POST("/login", handler.LoginUser)

	r.POST("/new", handler.NewReport)

	r.POST("/api/ai/detect", AI.DetectHazard)

	r.GET("/user", handler.GetReportsAndPoints)

	r.GET("/reports", handler.GetAllReports)
	r.GET("/reports/pending/", handler.GetPendingReports)
	r.GET("/reports/hospitals", handler.GetReportsForHospitals) // Added route for hospitals
    r.GET("/reports/police", handler.GetReportsForPolice)       // Added route for police
    r.GET("/reports/pwd", handler.GetReportsForPWD)             // Added route for PWD

	
	r.POST("/reports/updateStatus", handler.UpdateReportStatus)

	r.POST("/similarReports",handler.SimilarReports)

	r.GET("/me",handler.Me)

	r.Run(":6969")
}
