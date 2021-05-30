package intefaces

import (
	"context"

	"Sharykhin/rent-car/domain/consumer/models"
)

type (
	// ConsumerRepositoryInterface describes repository interface
	ConsumerRepositoryInterface interface {
		CreateConsumer(ctx context.Context, consumer *models.ConsumerModel) (*models.ConsumerModel, error)
	}
)
