package services

import (
	"context"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car"
	"Sharykhin/rent-car/domain/car/intefaces"
	"Sharykhin/rent-car/domain/car/types"
)

type (
	// CarService describes general business use-cases
	CarService struct {
		carRepo intefaces.CarRepositoryInterface
	}
)

// NewCarService is function constructor that returns a new instance of car service
func NewCarService(carRepo intefaces.CarRepositoryInterface) *CarService {
	srv := CarService{
		carRepo: carRepo,
	}

	return &srv
}

// CreateNewCar creates a new car
func (srv *CarService) CreateNewCar(ctx context.Context, m types.Model) (*car.CarModel, error) {
	c, err := car.NewCarModel(m)
	if err != nil {
		return nil, fmt.Errorf("[car][services][CreateNewCar] failed to create a new instance of a car model: %w", err)
	}

	c, err = srv.carRepo.Create(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("[car][services][CreateNewCar] failed to craete a new car: %w", err)
	}

	return c, nil
}

// GetCarByID returns a specific car by its ID
func (srv *CarService) GetCarByID(ctx context.Context, ID domain.ID) (*car.CarModel, error) {
	c, err := srv.carRepo.GetCarByID(ctx, ID)

	if err != nil {
		return nil, fmt.Errorf("[car][services][CreateNewCar] failed to get a car from the car service: %w", err)
	}

	return c, err
}
