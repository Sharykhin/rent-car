package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"Sharykhin/rent-car/api/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Environment variable SERVER_PORT is not defined")
	}

	web.Start(serverPort)
}
