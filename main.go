package main

import (
	"log"
	"wisewallet/database"
	"wisewallet/routes"
	"wisewallet/config"	
)

func main() {
	config.LoadConfig()

	database.RunMigration()

	database.Connect()

	r := routes.SetupRouter()

	log.Println("Server is running...")
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}