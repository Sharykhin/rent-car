package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
	"Sharykhin/rent-car/domain/car/value"
)

var (
	// ErrCarModelRequired describes error that model is required
	ErrCarModelRequired = errors.New("car model is required")
	// ErrCarModelInvalid
	ErrCarModelInvalid = errors.New("car model is invalid")
)

// IsCarModelCorrectSpecification checks whether car model is correct
func IsCarModelCorrectSpecification(car *model.CarModel) error {
	isCarModelEmpty := car.Model == ""
	if isCarModelEmpty {
		return domain.NewError(
			ErrCarModelRequired,
			"[domain][car][specification][IsCarModelCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	switch car.Model {
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
