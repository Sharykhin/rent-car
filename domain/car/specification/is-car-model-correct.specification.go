package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/value"
)

var (
	// ErrCarModelRequired describes error that model is required
	ErrCarModelRequired = errors.New("car model is required")
	// ErrCarModelInvalid means that car model is not from a valid list
	ErrCarModelInvalid = errors.New("car model is invalid")
)

// IsCarModelCorrectSpecification checks whether car model is correct
func IsCarModelCorrectSpecification(model value.Model) error {
	isCarModelEmpty := model == ""
	if isCarModelEmpty {
		return domain.NewError(
			ErrCarModelRequired,
			"[domain][car][specification][IsCarModelCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	switch model {
	case value.BMW:
	case value.Audi:
	default:
		return domain.NewError(
			ErrCarModelInvalid,
			"[domain][car][specification][IsCarModelCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	return nil
}
