package routes

import (
	"wisewallet/controllers"
	"wisewallet/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/auth/login", controllers.LoginUser)
	r.POST("/api/auth/register", controllers.RegisterUser)

	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/transactions", controllers.GetAllTransactions)
		api.POST("/transactions", controllers.CreateTransaction)
		api.GET("/transactions/:id", controllers.GetTransactionByID)
		api.PUT("/transactions/:id", controllers.UpdateTransaction)
		api.DELETE("/transactions/:id", controllers.DeleteTransaction)

		api.GET("/categories", controllers.GetAllCategories)
		api.POST("/categories", controllers.CreateCategory)
		api.GET("/categories/:id", controllers.GetCategoryByID)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)

		api.GET("/gamification", controllers.GetGamification)
		api.POST("/gamification/points", controllers.UpdatePoints)
		api.POST("/gamification/badges", controllers.AddBadge)
	}

	return r
}