package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()                   // gin router init
	router.GET("test", func(c *gin.Context) { // redirect page, http://localhost:8088/test
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	router.GET("/test1", func(c *gin.Context) { // redirect path, http://localhost:8088/test1
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
	})
	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
