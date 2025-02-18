package main

import (
	"log"

	"github.com/jevvonn/post-management-api/cmd/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.NewHTTPServer()
}
