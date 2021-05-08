package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"Sharykhin/rent-car/api/web"
)

func main() {
	l, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatalf("failed to parse log level env variable")
	}
	log.SetLevel(l)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("Environment variable SERVER_PORT is not defined")
	}

	web.Start(serverPort)
}
