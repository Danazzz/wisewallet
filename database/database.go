package database

import (
	"database/sql"
	"log"
	"wisewallet/config"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase(cfg *config.Config) {
	var err error
	DB, err = sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	log.Println("Connected to the database successfully")

	runMigrations()
}

func runMigrations() {
	migrationsDir := "./database/migrations"
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		filePath := filepath.Join(migrationsDir, file.Name())
		query, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", filePath, err)
		}

		_, err = DB.Exec(string(query))
		if err != nil {
			log.Fatalf("Failed to execute migration file %s: %v", filePath, err)
		}

		log.Printf("Executed migration: %s", file.Name())
	}
}