package main

import (
	"app/router"
)

func main() {
	router := router.SetRouterPublic()
	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
