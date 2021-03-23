package services

import (
	"context"
	"fmt"

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

func (s *CarService) CreateNewCar(ctx context.Context, model models.Model) (*models.Car, error) {
	car := models.NewCar(model)

	car, err := s.carRepository.Create(ctx, *car)
	if err != nil {
		return nil, fmt.Errorf("failed to craete a new car: %v", err)
	}

	return car, nil
}
