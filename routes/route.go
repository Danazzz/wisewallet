package routes

import (
	"github.com/gin-gonic/gin"
	"wisewallet/controllers"
	"wisewallet/middlewares"
	"wisewallet/config"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	auth := router.Group("/")
	auth.Use(middlewares.BasicAuthMiddleware())
	{
		auth.GET("/transactions", controllers.GetTransactions)
		auth.POST("/transactions", controllers.CreateTransaction)
		auth.GET("/transactions/:id", controllers.GetTransactionByID)
		auth.PUT("/transactions/:id", controllers.UpdateTransaction)
		auth.DELETE("/transactions/:id", controllers.DeleteTransaction)

		auth.GET("/badges", controllers.GetBadges)
		auth.POST("/badges", controllers.CreateBadge)
		auth.GET("/badges/:id", controllers.GetBadgeByID)
		auth.PUT("/badges/:id", controllers.UpdateBadge)
		auth.DELETE("/badges/:id", controllers.DeleteBadge)

		auth.POST("/user_badges", controllers.AddUserBadge)
		auth.GET("/user_badges", controllers.GetUserBadges)
		auth.DELETE("/user_badges/:id", controllers.DeleteUserBadge)
	}
}