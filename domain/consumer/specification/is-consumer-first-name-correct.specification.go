package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/model"
)

const (
	firstNameMaxLength = 50
)

var (
	ErrConsumerFirstNameRequired = errors.New("first name is empty")
	ErrConsumerFirstNameTooLong  = errors.New("first name is too long")
)

// IsConsumerFirstNameCorrectSpecification validates consumer last name
func IsConsumerFirstNameCorrectSpecification(consumer *model.ConsumerModel) error {
	isFirstNameEmpty := consumer.FirstName == ""
	if isFirstNameEmpty {
		return domain.NewError(
			ErrConsumerFirstNameRequired,
			"[domain][consumer][specification][IsConsumerFirstNameCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	if len(consumer.FirstName) > firstNameMaxLength {
		return domain.NewError(
			ErrConsumerFirstNameTooLong,
			"[domain][consumer][specification][IsConsumerFirstNameCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	return nil
}
