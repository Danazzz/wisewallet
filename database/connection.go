package database

import (
	"log"
	"wisewallet/config"
	"wisewallet/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	log.Println("Initializing database connection...")

	dsn := config.DatabaseURL
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully")
}

func RunMigration() {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	log.Println("Running database migrations...")
	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Transaction{},
		&models.Gamification{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully")
}