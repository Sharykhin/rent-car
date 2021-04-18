package services

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/intefaces"
	"Sharykhin/rent-car/domain/consumer/models"
	"context"
	"errors"
	"fmt"
)

type ConsumerService struct {
	consumerRepository intefaces.ConsumerRepositoryInterface
}

func NewConsumerService(consumerRepository intefaces.ConsumerRepositoryInterface) *ConsumerService {
	srv := ConsumerService{
		consumerRepository: consumerRepository,
	}

	return &srv
}

// CreateNewCar create a new car
func (s *ConsumerService) CreateNewConsumer(ctx context.Context, firstName, lastName, email string) (*models.Consumer, error) {
	consumer, err := models.NewConsumer(firstName, lastName, email, make([]models.Requisition, 0))
	if err != nil {
		return nil, domain.WrapError(errors.New("failed to create a new consumer model"), err)
	}

	consumer, err = s.consumerRepository.Create(ctx, *consumer)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new consumer in the consumer service: %w", err)
	}

	return consumer, nil
}
