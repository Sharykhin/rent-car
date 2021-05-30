package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
)

var (
	ErrConsumerLastNameRequired = errors.New("[consumer][IsConsumerLastNameCorrectSpecification] last name is empty")
	ErrConsumerLastNameTooLong  = errors.New("[consumer][IsConsumerLastNameCorrectSpecification] last name is too long")
)

// IsConsumerLastNameCorrectSpecification validates consumer first name
func IsConsumerLastNameCorrectSpecification(consumer *models.ConsumerModel) error {
	isEmpty := consumer.LastName == ""
	if isEmpty {
		return domain.NewError(ErrConsumerLastNameRequired, domain.ValidationErrorCode, "last name is required")
	}

	if len(consumer.LastName) > 50 {
		return domain.NewError(ErrConsumerLastNameTooLong, domain.ValidationErrorCode, "last name is too long")
	}

	return nil
}
