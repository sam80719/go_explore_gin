package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()

	// 32Mib
	//router.MaxMultipartMemory = 8 // limit memory

	router.Static("/", "./public") // http://localhost:8088/
	router.POST("/upload", func(c *gin.Context) {
		account := c.PostForm("account")
		password := c.PostForm("password")

		file, err := c.FormFile("file")

		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("error upload file: %s", err.Error()))
			return
		}

		fileName := filepath.Base(file.Filename)

		if err := c.SaveUploadedFile(file, fileName); err != nil {
			c.String(http.StatusOK, fmt.Sprintf("error save file: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("account=%s password=%s upload file name is %s", account, password, fileName))
	})

	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
