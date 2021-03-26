package intefaces

import (
	"Sharykhin/rent-car/domain"
	"context"

	"Sharykhin/rent-car/domain/car/models"
)

type (
	CarRepositoryInterface interface {
		Create(ctx context.Context, car models.Car) (*models.Car, error)
		GetCarByID(ctx context.Context, ID domain.ID) (*models.Car, error)
	}
)
