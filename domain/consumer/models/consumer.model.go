package models

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
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

func (cs *Consumer) RentCar(c *models.CarModel) error {
	if len(cs.Requisitions) > 2 {
		return domain.RequisitionLimitExceededError
	}

	requisition := NewRequisition(c)
	cs.Requisitions = append(cs.Requisitions, *requisition)

	return nil
}

// NewConsumer creates a new consumer
func NewConsumer(firstName, lastName, email string, requisitions []Requisition) (*Consumer, error) {
	// TODO: Move all validation rules under specifications
	if len(firstName) == 0 || len(firstName) > 50 {
		return nil, domain.NewError(errors.New("consumer first name is invalid"), domain.ValidationErrorCode, "First name is invalid.")
	}

	if len(lastName) == 0 || len(firstName) > 50 {
		return nil, domain.NewError(errors.New("consumer last name is invalid"), domain.ValidationErrorCode, "Last name is invalid.")
	}

	if len(email) == 0 || len(firstName) > 80 {
		return nil, domain.NewError(errors.New("consumer email is invalid"), domain.ValidationErrorCode, "Email is invalid.")
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
