package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default() // gin router init
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hallo %s", name)
	})

	err := router.Run(":8088") // web path：http://localhost:8088/
	if err != nil {
		panic(err)
	}
}
