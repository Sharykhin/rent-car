package service

import (
	"context"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
)

type (
	// CarRepositoryInterface describes car repository methods
	CarRepositoryInterface interface {
		CreateCar(ctx context.Context, c *model.CarModel) (*model.CarModel, error)
		GetCarByID(ctx context.Context, ID domain.ID) (*model.CarModel, error)
		UpdateCar(ctx context.Context, car *model.CarModel) error
	}
	// FileStorageInterface describes api of storing files
	FileStorageInterface interface {
		Upload(ctx context.Context, path string, data []byte) error
	}
)
