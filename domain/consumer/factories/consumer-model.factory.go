package factories

import (
	"fmt"
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
	"Sharykhin/rent-car/domain/consumer/specification"
)

// NewConsumerModel creates a new consumer
func NewConsumerModel(firstName, lastName, email string, requisitions []models.Requisition) (*models.ConsumerModel, error) {

	consumer := models.ConsumerModel{
		ID:           domain.Empty(),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		CreatedAt:    time.Now().UTC(),
		Requisitions: requisitions,
	}

	err := specification.IsConsumerFirstNameCorrectSpecification(&consumer)
	if err != nil {
		return nil, fmt.Errorf("[consumer][NewConsumerModel] failed to pass specification: %w", err)
	}

	err = specification.IsConsumerLastNameCorrectSpecification(&consumer)
	if err != nil {
		return nil, fmt.Errorf("[consumer][NewConsumerModel] failed to pass specification: %w", err)
	}

	err = specification.IsConsumerEmailCorrectSpecification(&consumer)
	if err != nil {
		return nil, fmt.Errorf("[consumer][NewConsumerModel] failed to pass specification: %w", err)
	}

	return &consumer, nil
}
