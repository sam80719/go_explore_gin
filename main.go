package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginForm struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.Static("/", "./public") // http://localhost:8088/
	router.POST("/login", func(c *gin.Context) {
		var form loginForm

		if err := c.ShouldBind(&form); err != nil { // binding value
			c.String(http.StatusUnauthorized, fmt.Sprintf("binding form failure: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("binding form success: %s", form.Account))
	})

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
