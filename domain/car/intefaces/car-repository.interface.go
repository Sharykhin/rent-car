package intefaces

import (
	"context"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car"
)

type (
	// CarRepositoryInterface represents car repository methods
	CarRepositoryInterface interface {
		Create(ctx context.Context, c *car.CarModel) (*car.CarModel, error)
		GetCarByID(ctx context.Context, ID domain.ID) (*car.CarModel, error)
	}
)
