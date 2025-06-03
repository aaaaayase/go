package router

import (
	"first-project/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	router := gin.Default()
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controller.Register)

		auth.POST("/register", controller.Login)
	}

	return router
}
