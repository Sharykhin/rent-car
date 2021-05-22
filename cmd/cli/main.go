package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"Sharykhin/rent-car/di"
)

// TODO: This is used for simple tests. Remove it at the end
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	container, err := di.Init()
	if err != nil {
		log.Fatalf("failed to initialize di container: %v", err)
	}
	postgres := container.PostgresConn
	err = postgres.Connect()

	fmt.Println(err)
}
