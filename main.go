package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()                      // gin router init
	router.GET("/", func(context *gin.Context) { // create root test
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
	})

	err := router.Run(":8088") // web path：http://localhost:8088/
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
