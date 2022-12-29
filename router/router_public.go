package router

import (
	"app/handler"
	"app/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouterPublic() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(), middleware.Auth(), gin.Recovery())

	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", handler.UserLogin)       // http://localhost:8088/user/login
		userRouter.POST("/register", handler.UserRegister) // http://localhost:8088/user/register

	}

	return router
}
