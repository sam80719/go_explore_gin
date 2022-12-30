package main

import (
	_ "app/docs" // swagger use
	"app/router"
)

// @title Gin swagger test
// @version 1.0
// @description Gin swagger ini

// @contact.name Sam Chen
// @contact.url https://github.com/sam80719

// @host localhost:8088
func main() {
	router := router.SetRouterPublic()
	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
