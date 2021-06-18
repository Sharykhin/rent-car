package service

import (
	"context"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/factory"
	"Sharykhin/rent-car/domain/consumer/model"
)

type (
	// ConsumerService describes business use-cases around consumer domain
	ConsumerService struct {
		consumerRepo ConsumerRepositoryInterface
	}
)

// NewConsumerService is a function constructor that creates a new instance of ConsumerService struct
func NewConsumerService(consumerRepo ConsumerRepositoryInterface) *ConsumerService {
	srv := ConsumerService{
		consumerRepo: consumerRepo,
	}

	return &srv
}

// CreateNewConsumer creates a new consumer
func (srv *ConsumerService) CreateNewConsumer(ctx context.Context, firstName, lastName, email string) (*model.ConsumerModel, error) {
	consumer, err := factories.NewConsumerModel(firstName, lastName, email)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][consumer][service][ConsumerService][CreateNewConsumer][NewConsumerModel]")
	}

	consumer, err = srv.consumerRepo.CreateConsumer(ctx, consumer)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][consumer][service][ConsumerService][CreateNewConsumer][CreateConsumer]")
	}

	return consumer, nil
}
