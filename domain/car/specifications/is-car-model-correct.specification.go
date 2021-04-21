package specifications

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car"
)

func NewIsCarModelCorrectSpecification(car *car.CarModel) error {
	if car.Model == "" {
		return domain.NewError(
			errors.New("[car][specifications][NewIsCarModelCorrectSpecification] car model is required"),
			domain.ValidationErrorCode,
			"Car model is required.",
		)
	}

	return nil
}
