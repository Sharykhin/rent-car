package main

import (
	"Sharykhin/rent-car/infrastructure/postgres"
	postgresRepositories "Sharykhin/rent-car/infrastructure/postgres/repositories"
	"fmt"
	"log"
	"os"

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
	car1 := carModels.NewCar(carModels.Audi)
	car2, err := carRepository.Create(car1)

	fmt.Println(car1, car2, err)
}
