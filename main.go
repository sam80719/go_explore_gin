package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/", "./public") // http://localhost:8088/
	router.POST("/login", func(c *gin.Context) {
		account := c.PostForm("account")
		password := c.PostForm("password")
		c.String(http.StatusOK, fmt.Sprintf("account=%s password=%s", account, password))
	})

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
