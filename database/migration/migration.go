package main

import (
	"fmt"
	"log"

	"github.com/jevvonn/post-management-api/database"
	"github.com/jevvonn/post-management-api/internal/model/domain"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDatabase()
	migrator := database.DB.Migrator()

	fmt.Println("Running migration...")
	migrator.AutoMigrate(domain.Post{})
	fmt.Println("Migration completed")
}
