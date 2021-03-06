package main

import (
	"Sharykhin/rent-car/logger"
	"log"
	"os"

	"github.com/joho/godotenv"

	"Sharykhin/rent-car/api/web"
	"Sharykhin/rent-car/di"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[web][main] failed to load .env file: %v", err)
	}

	err = logger.Init()
	if err != nil {
		log.Fatalf("[web][main] failed to initialize logger: %v", err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("[web][main] environment variable SERVER_PORT is not defined")
	}

	err = di.Init()
	if err != nil {
		log.Fatalf("[web][main] failed to initialize di: %v", err)
	}

	server := web.NewServer(serverPort)

	server.ListenAndServe()
}
