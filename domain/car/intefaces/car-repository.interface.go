package intefaces

import (
	"context"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
)

type (
	// CarRepositoryInterface describes car repository methods
	CarRepositoryInterface interface {
		CreateCar(ctx context.Context, c *models.CarModel) (*models.CarModel, error)
		GetCarByID(ctx context.Context, ID domain.ID) (*models.CarModel, error)
	}
)
