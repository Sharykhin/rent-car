package service

import (
	"context"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/factory"
	"Sharykhin/rent-car/domain/car/intefaces"
	"Sharykhin/rent-car/domain/car/models"
	"Sharykhin/rent-car/domain/car/types"
)

type (
	// CarService describes general business use-cases around car domain
	CarService struct {
		carRepo intefaces.CarRepositoryInterface
	}
)

// NewCarService creates a new car service instance
func NewCarService(carRepo intefaces.CarRepositoryInterface) *CarService {
	srv := CarService{
		carRepo: carRepo,
	}

	return &srv
}

// CreateNewCar creates a new car
func (srv *CarService) CreateNewCar(ctx context.Context, model types.Model) (*models.CarModel, error) {
	car, err := factory.NewCarModel(model)
	if err != nil {
		return nil, fmt.Errorf("[domain][car][CarService][CreateNewCar] failed to create a new car model: %w", err)
	}

	car, err = srv.carRepo.CreateCar(ctx, car)
	if err != nil {
		return nil, fmt.Errorf("[domain][car][CarService][CreateNewCar] repository failed to craete a new car: %w", err)
	}

	return car, nil
}

// GetCarByID returns a specific car by its ID
func (srv *CarService) GetCarByID(ctx context.Context, ID domain.ID) (*models.CarModel, error) {
	c, err := srv.carRepo.GetCarByID(ctx, ID)

	if err != nil {
		return nil, fmt.Errorf("[CarService][GetCarByID] failed to get a car from the car repository: %w", err)
	}

	return c, err
}
