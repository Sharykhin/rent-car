package specification

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
)

func TestIsConsumerFirstNameCorrectSpecification(t *testing.T) {
	faker := gofakeit.New(0)
	assert := assert.New(t)

	consumer := models.ConsumerModel{
		ID:           domain.Empty(),
		LastName:     faker.LastName(),
		Email:        faker.Email(),
		Requisitions: make([]models.Requisition, 0),
	}

	tt := []struct {
		name          string
		inFirstName   string
		expectedError error
	}{
		{
			name:          "First name is required",
			inFirstName:   "",
			expectedError: ErrConsumerFirstNameRequired,
		},
		{
			name:          "Fist name is too long",
			inFirstName:   faker.LetterN(51),
			expectedError: ErrConsumerFirstNameTooLong,
		},
		{
			name:          "First name is valid",
			inFirstName:   faker.FirstName(),
			expectedError: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			consumer.FirstName = tc.inFirstName
			err := IsConsumerFirstNameCorrectSpecification(&consumer)
			assert.True(errors.Is(err, tc.expectedError))
		})
	}
}
