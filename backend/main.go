package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/handler"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8000"}, // Change this if your frontend URL changes
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", handler.GetData)

	r.POST("/signup", handler.RegisterUser)

	r.POST("/login", handler.LoginUser)

	r.POST("/new", handler.RoleMiddleware("user"), handler.NewReport)

	r.GET("/user", handler.GetReportsAndPoints)

	r.GET("/reports", handler.RoleMiddleware("admin"), handler.GetAllReports)
	r.GET("/reports/pending/", handler.GetPendingReports)
	r.GET("/reports/hospitals", handler.RoleMiddleware("hospital"), handler.GetReportsForHospitals) // route for hospitals
	r.GET("/reports/police", handler.RoleMiddleware("police"), handler.GetReportsForPolice)         // route for police
	r.GET("/reports/pwd", handler.RoleMiddleware("pwd"), handler.GetReportsForPWD)                  // route for PWD

	r.POST("/reports/updateStatus", handler.RoleMiddleware("hospital", "pwd", "police", "admin"), handler.UpdateReportStatus)

	r.POST("/similarReports", handler.SimilarReports)

	r.GET("/me", handler.Me)

	r.Run(":6969")
}
