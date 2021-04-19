package services

import (
	"context"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/intefaces"
	"Sharykhin/rent-car/domain/car/models"
)

type CarService struct {
	carRepo intefaces.CarRepositoryInterface
}

func NewCarService(carRepo intefaces.CarRepositoryInterface) *CarService {
	srv := CarService{
		carRepo: carRepo,
	}

	return &srv
}

// CreateNewCar creates a new car
func (srv *CarService) CreateNewCar(ctx context.Context, model models.Model) (*models.Car, error) {
	car, err := models.NewCar(model)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new instance of car model: %w", err)
	}

	car, err = srv.carRepo.Create(ctx, *car)
	if err != nil {
		return nil, fmt.Errorf("failed to craete a new car: %w", err)
	}

	return car, nil
}

// GetCarByID returns a specific car by its ID
func (srv *CarService) GetCarByID(ctx context.Context, ID domain.ID) (*models.Car, error) {
	car, err := srv.carRepo.GetCarByID(ctx, ID)

	if err != nil {
		return nil, fmt.Errorf("failed to get a car from the car service: %w", err)
	}

	return car, err
}
