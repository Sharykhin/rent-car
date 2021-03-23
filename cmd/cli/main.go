package main

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/services"
	"Sharykhin/rent-car/infrastructure/postgres"
	postgresRepositories "Sharykhin/rent-car/infrastructure/postgres/repositories"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	carModels "Sharykhin/rent-car/domain/car/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresURL := os.Getenv("POSTGRES_URL")
	db, err := postgres.Connect(postgresURL)

	carRepository := postgresRepositories.NewCarRepository(db)
	carService := services.NewCarService(carRepository)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	car2, err := carService.CreateNewCar(ctx, carModels.BMW)
	if err != nil {
		fmt.Println(errors.Is(err, domain.InvalidCarModelError))
	}

	fmt.Println(car2, err)
}
