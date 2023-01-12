package router

import (
	"app/handler"
	"app/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouterPublic() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(), middleware.Auth(), gin.Recovery())

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:id", handler.GetOne)     // http://localhost:8088/user/:id
		userRouter.GET("/all", handler.UserGetAll) // http://localhost:8088/user/all
		userRouter.DELETE("/:id", handler.DeleteOne)
		userRouter.POST("/login", handler.UserLogin)       // http://localhost:8088/user/login
		userRouter.POST("/register", handler.UserRegister) // http://localhost:8088/user/register
	}

	url := ginSwagger.URL("http://localhost:8088/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url)) // http://localhost:8088/swagger/index.html

	return router
}
