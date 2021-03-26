package models

import (
	"Sharykhin/rent-car/domain"
	carModels "Sharykhin/rent-car/domain/car/models"
	"errors"
	"time"
)

// Consumer is a person/organization who/what rents cars
type Consumer struct {
	ID           domain.ID
	FirstName    string
	LastName     string
	Email        string
	CreatedAt    time.Time
	Requisitions []Requisition
}

func (c *Consumer) RentCar(car carModels.Car) error {
	if len(c.Requisitions) > 2 {
		return domain.RequisitionLimitExceededError
	}

	requisition := NewRequisition(car)
	c.Requisitions = append(c.Requisitions, *requisition)

	return nil
}

// NewConsumer creates a new consumer
func NewConsumer(firstName, lastName, email string, requisitions []Requisition) (*Consumer, error) {
	// TODO: Move all validation rules under specifications
	if len(firstName) == 0 || len(firstName) > 50 {
		return nil, domain.NewError(errors.New("consumer first name is invalid"), domain.ValidationErrorCode)
	}

	if len(lastName) == 0 || len(firstName) > 50 {
		return nil, domain.NewError(errors.New("consumer last name is invalid"), domain.ValidationErrorCode)
	}

	if len(email) == 0 || len(firstName) > 80 {
		return nil, domain.NewError(errors.New("consumer email is invalid"), domain.ValidationErrorCode)
	}

	consumer := Consumer{
		ID:           domain.Empty(),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		CreatedAt:    time.Now().UTC(),
		Requisitions: requisitions,
	}

	return &consumer, nil
}
