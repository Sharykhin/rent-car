package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"Sharykhin/rent-car/api/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[web][main] failed to load .env file: %v", err)
	}

	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatalf("[web][main] failed to parse a log level env variable: %v", err)
	}
	log.SetLevel(level)
	logger := log.WithField("service", os.Getenv("SERVICE_ID"))

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("[web][main] environment variable SERVER_PORT is not defined")
	}

	server := web.NewServer(serverPort, logger)

	server.ListenAndServe()
}
