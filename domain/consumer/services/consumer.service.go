package services

import (
	"context"
	"fmt"

	"Sharykhin/rent-car/domain/consumer/factories"
	"Sharykhin/rent-car/domain/consumer/intefaces"
	"Sharykhin/rent-car/domain/consumer/models"
)

type (
	// ConsumerService describes business use-cases around consumer domain
	ConsumerService struct {
		consumerRepo intefaces.ConsumerRepositoryInterface
	}
)

// NewConsumerService is a function constructor that creates a new instance of ConsumerService struct
func NewConsumerService(consumerRepo intefaces.ConsumerRepositoryInterface) *ConsumerService {
	srv := ConsumerService{
		consumerRepo: consumerRepo,
	}

	return &srv
}

// CreateNewConsumer creates a new consumer
func (srv *ConsumerService) CreateNewConsumer(ctx context.Context, firstName, lastName, email string) (*models.ConsumerModel, error) {
	consumer, err := factories.NewConsumerModel(firstName, lastName, email, make([]models.Requisition, 0))
	if err != nil {
		return nil, fmt.Errorf("[ConsumerService][CreateNewConsumer] failed to craete a new consumer model: %w", err)
	}

	consumer, err = srv.consumerRepo.CreateConsumer(ctx, consumer)
	if err != nil {
		return nil, fmt.Errorf("[ConsumerService][CreateNewConsumer] repository failed to craete a new consumer: %w", err)
	}

	return consumer, nil
}
