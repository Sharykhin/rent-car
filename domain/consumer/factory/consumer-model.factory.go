package factories

import (
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/model"
	"Sharykhin/rent-car/domain/consumer/specification"
)

// NewConsumerModel creates a new consumer including all validation steps
func NewConsumerModel(firstName, lastName, email string) (*model.ConsumerModel, error) {

	consumer := model.ConsumerModel{
		ID:        domain.Empty(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}

	err := specification.IsConsumerFirstNameCorrectSpecification(&consumer)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][consumer][factory][NewConsumerModel]")
	}

	err = specification.IsConsumerLastNameCorrectSpecification(&consumer)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][consumer][factory][NewConsumerModel]")
	}

	err = specification.IsConsumerEmailCorrectSpecification(&consumer)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][consumer][factory][NewConsumerModel]")
	}

	return &consumer, nil
}
