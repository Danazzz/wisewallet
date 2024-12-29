package main

import (
	"wisewallet/config"
	"wisewallet/database"
	"wisewallet/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg)

	router := gin.Default()

	routes.SetupRoutes(router, cfg)

	router.Run(":8080")
}