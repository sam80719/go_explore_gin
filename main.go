package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()                              // gin router init
	router.GET("/testJSON", func(context *gin.Context) { // http://localhost:8088/testJSON
		context.JSON(http.StatusOK, gin.H{"message": "test json", "status": "OK"})
	})
	router.GET("/testJSON1", func(context *gin.Context) { // http://localhost:8088/testJSON1
		var message struct {
			Name  string
			Sex   string
			Phone int
		}

		message.Name = "Sam"
		message.Sex = "man"
		message.Phone = 95279527
		context.JSON(http.StatusOK, message)
	})
	// follow above, we knew json can support struct

	router.GET("/testXML", func(context *gin.Context) { // http://localhost:8088/testXML
		context.XML(http.StatusOK, gin.H{"message": "test XML", "status": "OK"})
	})

	router.GET("/testYAML", func(context *gin.Context) { // http://localhost:8088/testYAML
		context.YAML(http.StatusOK, gin.H{"message": "test YAML", "status": "OK"})
	})

	router.GET("/testHTML", func(context *gin.Context) { // http://localhost:8088/testHTML
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<h1>HELLO WORLD!!!</h1>",
		})
	})
	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
