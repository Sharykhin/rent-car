package main

import (
	"Sharykhin/rent-car/infrastructure/postgres/query"
	"context"
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

	err = di.Init()
	if err != nil {
		log.Fatalf("failed to initialize di container: %v", err)
	}
	postgres := di.Container.PostgresConn
	err = postgres.Connect()
	defer postgres.Close()

	carQueryRepository := query.NewPostgresCarQuery(postgres)
	cars, total, err := carQueryRepository.GetPagedCarsList(context.TODO(), 2, 0)

	fmt.Println(cars, total, err)
}
