package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/handler"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			// Allow Codespaces frontend and localhost
			return strings.HasSuffix(origin, ".app.github.dev") || origin == "http://localhost:5173"
		},
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


	//admin routes
	//note: remove the users role from the middleware in production
	r.GET("admin/users/unverified", handler.RoleMiddleware("admin","user"), handler.GetUnverifiedUsers)
	r.POST("/admin/users/verify", handler.RoleMiddleware("admin","user"), handler.VerifyUser)


	r.GET("/me", handler.Me)

	r.Run(":6969")
}
