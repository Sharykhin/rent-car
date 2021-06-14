package main

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/factory"
	"Sharykhin/rent-car/domain/car/specification"
	"Sharykhin/rent-car/infrastructure/postgres/query"
	"Sharykhin/rent-car/infrastructure/postgres/repositories"
	"context"
	"errors"
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

	car, err := factory.NewCarModel("")

	fmt.Println(err, car, errors.Is(err, specification.ErrCarModelRequired))

	car, err = di.Container.PostgresCarRepository.GetCarByID(context.Background(), domain.ID("d92b94c8-6d3f-4663-b5da-f61c653eb898"))
	fmt.Println(car, err, errors.Is(err, repositories.ErrCarNotFound))

}
