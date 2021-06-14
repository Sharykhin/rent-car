package specification

import (
	"errors"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
)

var (
	// ErrCarModelRequired describes error that model is required
	ErrCarModelRequired = errors.New("car model is required")
)

// IsCarModelCorrectSpecification checks whether car model is correct
func IsCarModelCorrectSpecification(car *model.CarModel) error {
	isCarModelEmpty := car.Model == ""
	if isCarModelEmpty {
		return domain.NewError(
			fmt.Errorf("[domain][car][specification][IsCarModelCorrectSpecification] %w", ErrCarModelRequired),
			domain.ValidationErrorCode,
			"Car model is required.",
		)
	}

	return nil
}
