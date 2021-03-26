package services

import (
	"context"
	"errors"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/intefaces"
	"Sharykhin/rent-car/domain/car/models"
)

type CarService struct {
	carRepository intefaces.CarRepositoryInterface
}

func NewCarService(carRepository intefaces.CarRepositoryInterface) *CarService {
	srv := CarService{
		carRepository: carRepository,
	}

	return &srv
}

// CreateNewCar create a new car
func (s *CarService) CreateNewCar(ctx context.Context, model models.Model) (*models.Car, error) {
	car := models.NewCar(model)

	car, err := s.carRepository.Create(ctx, *car)
	if err != nil {
		return nil, fmt.Errorf("failed to craete a new car: %v", err)
	}

	return car, nil
}

// GetCarByID returns a specific car by its ID
func (s *CarService) GetCarByID(ctx context.Context, ID domain.ID) (*models.Car, error) {
	car, err := s.carRepository.GetCarByID(ctx, ID)

	if err != nil {
		return nil, domain.WrapError(errors.New("failed to get a car from the car service"), err)
	}

	return car, err
}
