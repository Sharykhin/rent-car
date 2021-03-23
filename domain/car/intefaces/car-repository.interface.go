package intefaces

import (
	"context"

	"Sharykhin/rent-car/domain/car/models"
)

type (
	CarRepositoryInterface interface {
		Create(ctx context.Context, car models.Car) (*models.Car, error)
	}
)
