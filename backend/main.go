package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/handler"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.GetData)

	r.POST("/signup", handler.RegisterUser)
	r.POST("/login", handler.LoginUser)

	r.Run(":6969")
}
