package specification

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
)

func TestIsConsumerLastNameCorrectSpecification(t *testing.T) {
	faker := gofakeit.New(0)
	assert := assert.New(t)

	consumer := models.ConsumerModel{
		ID:           domain.Empty(),
		FirstName:    faker.FirstName(),
		Email:        faker.Email(),
		Requisitions: make([]models.Requisition, 0),
	}

	tt := []struct {
		name          string
		inLastName    string
		expectedError error
	}{
		{
			name:          "Last name is required",
			inLastName:    "",
			expectedError: ErrConsumerLastNameRequired,
		},
		{
			name:          "Last name is too long",
			inLastName:    faker.LetterN(51),
			expectedError: ErrConsumerLastNameTooLong,
		},
		{
			name:          "Last name is valid",
			inLastName:    faker.LastName(),
			expectedError: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			consumer.LastName = tc.inLastName
			err := IsConsumerLastNameCorrectSpecification(&consumer)
			assert.True(errors.Is(err, tc.expectedError))
		})
	}
}
