package main

import (
	"log"
	"wisewallet/config"
	"wisewallet/database"
	"wisewallet/routes"
)

func main() {
	config.LoadConfig()

	database.Init()

	database.RunMigration()

	r := routes.SetupRouter()

	log.Println("Server is running...")
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}