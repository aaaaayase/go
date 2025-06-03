package router

import (
	"first-project/controller"
	"first-project/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	router := gin.Default()
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controller.Register)
		auth.POST("/register", controller.Login)

	}

	api := router.Group("/api")
	api.GET("/exchangeRates", controller.GetExchangeRates)
	api.Use(middleware.AuthMiddleWare())
	{
		api.POST("/exchangeRates", controller.CreateExchangeRate)
		api.POST("/articles", controller.CreateArticle)
		api.GET("/articles", controller.GetArticle)
		api.GET("/articles/:id", controller.GetArticleByID)

		api.POST("/articles/:id/like", controller.LikeArticle)
		api.GET("/articles/:id/like", controller.GetArticleLikes)
	}

	return router
}
