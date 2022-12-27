package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*") // load temple file under directory view

	//router.LoadHTMLFiles("view/index.tmpl", "view/dashBoard.tmpl") // you also can use  LoadHTMLFiles, but we need write existing file

	router.GET("/index", func(c *gin.Context) { // http://localhost:8088/index
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "learn Golang",
			"message": "hello world!",
			"name":    "sam",
		})
	})

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
