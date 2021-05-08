package factories

import (
	"errors"
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
)

// NewConsumerModel creates a new consumer
func NewConsumerModel(firstName, lastName, email string, requisitions []models.Requisition) (*models.ConsumerModel, error) {

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

	consumer := models.ConsumerModel{
		ID:           domain.Empty(),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		CreatedAt:    time.Now().UTC(),
		Requisitions: requisitions,
	}

	return &consumer, nil
}
