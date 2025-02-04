package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leoantony72/Uyir/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Some fields are missing", "err": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error hashing password", "err": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	userKey := "user:" + user.Name

	userData, err := json.Marshal(user)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error marshaling user data", "err": err.Error()})
		return
	}

	r := Db.Set(Ctx, userKey, userData, 0)
	fmt.Println(r)

	c.JSON(200, gin.H{"message": "User registered"})
}
