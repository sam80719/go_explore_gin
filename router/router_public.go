package router

import (
	"app/handler"
	"github.com/gin-gonic/gin"
)

func SetRouterPublic() *gin.Engine { // func name use Upper Camal Case can be used by other file
	router := gin.Default() // gin router init

	userRouter := router.Group("/user")
	{ // <== code style
		userRouter.GET("/:name", handler.UserSave)   // http://localhost:8088/user/sam
		userRouter.GET("/", handler.UserSaveBYQuery) // http://localhost:8088/user/?name=sam, http://localhost:8088/user/?name=sam&phone=85918591

	}

	return router
}
