package service

import (
	"context"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/factory"
	"Sharykhin/rent-car/domain/car/model"
	"Sharykhin/rent-car/domain/car/value"
)

type (
	// CarService describes general business use-cases around car domain
	CarService struct {
		carRepo CarRepositoryInterface
	}
)

// NewCarService creates a new car service instance
func NewCarService(carRepo CarRepositoryInterface) *CarService {
	srv := CarService{
		carRepo: carRepo,
	}

	return &srv
}

// CreateNewCar creates a new car
func (srv *CarService) CreateNewCar(ctx context.Context, model value.Model) (*model.CarModel, error) {
	car, err := factory.NewCarModel(model)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][CarService][CreateNewCar]")
	}

	car, err = srv.carRepo.CreateCar(ctx, car)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][CarService][CreateNewCar]")
	}

	return car, nil
}

// GetCarByID returns a specific car by its ID
func (srv *CarService) GetCarByID(ctx context.Context, ID domain.ID) (*model.CarModel, error) {
	c, err := srv.carRepo.GetCarByID(ctx, ID)

	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][CarService][GetCarByID]")
	}

	return c, err
}
