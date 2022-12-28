package router

import (
	"app/handler"
	"github.com/gin-gonic/gin"
)

func SetRouterPublic() *gin.Engine {
	router := gin.Default() // gin router init

	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", handler.UserLogin) // http://localhost:8088/user/login

	}

	return router
}
