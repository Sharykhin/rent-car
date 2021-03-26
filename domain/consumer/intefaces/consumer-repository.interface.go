package intefaces

import (
	"context"

	"Sharykhin/rent-car/domain/consumer/models"
)

type (
	ConsumerRepositoryInterface interface {
		Create(ctx context.Context, consumer models.Consumer) (*models.Consumer, error)
	}
)
