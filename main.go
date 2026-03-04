package main

import (
	"log"

	"main/kudoserver"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables at application startup
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	if err := kudoserver.StartServer(":8080"); err != nil {
		log.Fatal(err)
	}
}
