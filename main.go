package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()             // gin router init
	router.POST("/login", useStructTag) // use cmd: curl -v --form account=account1 --form password=passwordd http://localhost:8088/login

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}

func useStructTag(c *gin.Context) {
	var f LoginForm
	account := c.PostForm("account") // select post form data
	fmt.Printf("account form is: %s\n", account)

	if c.ShouldBind(&f) == nil { // by form
		if f.Account == "account1" && f.Password == "password1" {
			c.JSON(http.StatusOK, gin.H{"message": "already login"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "login failure"})
		}
	}
}
