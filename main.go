package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	t := time.Now()
	//formatted := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	formatted := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day()) // time format
	file, _ := os.Create(formatted + "gin.log")

	gin.DefaultWriter = io.MultiWriter(file)

	router := gin.Default()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string { // record more detail
		return fmt.Sprintf("format log: %s - [%s] %s %s", params.ClientIP, params.TimeStamp, params.Method, params.Request)
	}))
	router.GET("/testLog", func(c *gin.Context) { // http://localhost:8088/testLog
		c.String(http.StatusOK, "hello world!")
	})

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
