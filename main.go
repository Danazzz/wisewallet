package main

import (
	"wisewallet/database"
	"wisewallet/routes"
	"wisewallet/config"	
)

func main() {
	config.LoadConfig()

	database.RunMigration()

	database.Connect()

	r := routes.SetupRouter()

	r.Run(":8080")
}