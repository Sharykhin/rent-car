package service

import (
	"Sharykhin/rent-car/domain"
	"context"

	"Sharykhin/rent-car/domain/consumer/model"
)

type (
	// ConsumerRepositoryInterface describes repository interface
	ConsumerRepositoryInterface interface {
		CreateConsumer(ctx context.Context, consumer *model.ConsumerModel) (*model.ConsumerModel, error)
		GetConsumerByID(ctx context.Context, ID domain.ID) (*model.ConsumerModel, error)
	}
	// FileStorageInterface describes api of storing files
	FileStorageInterface interface {
		Upload(ctx context.Context, path string, data []byte) error
	}
)
