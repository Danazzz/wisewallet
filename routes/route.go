package routes

import (
	"github.com/gin-gonic/gin"
	"wisewallet/controllers"
	"wisewallet/middlewares"
	"wisewallet/config"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Auth Routes
	router.POST("/api/auth/register", controllers.RegisterUser)
	router.POST("/api/auth/login", controllers.LoginUser)

	// Grouping routes with Basic Auth middleware
	auth := router.Group("/api")
	auth.Use(middlewares.BasicAuthMiddleware())
	{
		// Transaction Routes
		auth.GET("/transactions", controllers.GetTransactions)
		auth.POST("/transactions", controllers.CreateTransaction)
		auth.GET("/transactions/:id", controllers.GetTransactionByID)
		auth.PUT("/transactions/:id", controllers.UpdateTransaction)
		auth.DELETE("/transactions/:id", controllers.DeleteTransaction)

		// Badge Routes
		auth.GET("/badges", controllers.GetBadges)
		auth.POST("/badges", controllers.CreateBadge)
		auth.GET("/badges/:id", controllers.GetBadgeByID)
		auth.PUT("/badges/:id", controllers.UpdateBadge)
		auth.DELETE("/badges/:id", controllers.DeleteBadge)

		// UserBadge Routes
		auth.POST("/user_badges", controllers.AddUserBadge)
		auth.GET("/user_badges", controllers.GetUserBadges)
		auth.GET("/user_badges/:id", controllers.GetUserBadgeByID)
		auth.DELETE("/user_badges/:id", controllers.DeleteUserBadge)
	}
}