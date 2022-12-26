package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func main() {
	router := gin.Default()           // gin router init
	router.GET("/user", useStructTag) // http://localhost:8088/user?name=Sam&email=sam80719@gmail.com

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}

func useStructTag(c *gin.Context) {
	var u User
	if c.ShouldBindQuery(&u) == nil {
		fmt.Println(u.Name)
		fmt.Println(u.Email)
	}

	c.String(http.StatusOK, "OK")
}
