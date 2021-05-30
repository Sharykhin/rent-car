package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
)

var (
	ErrConsumerFirstNameRequired = errors.New("[consumer][IsConsumerFirstNameCorrectSpecification] first name is empty")
	ErrConsumerFirstNameTooLong  = errors.New("[consumer][IsConsumerFirstNameCorrectSpecification] first name is too long")
)

// IsConsumerFirstNameCorrectSpecification validates consumer last name
func IsConsumerFirstNameCorrectSpecification(consumer *models.ConsumerModel) error {
	isEmpty := consumer.FirstName == ""
	if isEmpty {
		return domain.NewError(ErrConsumerFirstNameRequired, domain.ValidationErrorCode, "first name is required")
	}

	if len(consumer.FirstName) > 50 {
		return domain.NewError(ErrConsumerFirstNameTooLong, domain.ValidationErrorCode, "first name is too long")
	}

	return nil
}
