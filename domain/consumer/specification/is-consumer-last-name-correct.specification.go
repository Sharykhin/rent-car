package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
)

const (
	lastNameMaxLength = 50
)

var (
	ErrConsumerLastNameRequired = errors.New("last name is empty")
	ErrConsumerLastNameTooLong  = errors.New("last name is too long")
)

// IsConsumerLastNameCorrectSpecification validates consumer last name
func IsConsumerLastNameCorrectSpecification(consumer *models.ConsumerModel) error {
	isEmpty := consumer.LastName == ""
	if isEmpty {
		return domain.NewError(
			ErrConsumerLastNameRequired,
			"[domain][consumer][specification][IsConsumerLastNameCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	if len(consumer.LastName) > lastNameMaxLength {
		return domain.NewError(
			ErrConsumerLastNameTooLong,
			"[domain][consumer][specification][IsConsumerLastNameCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	return nil
}
