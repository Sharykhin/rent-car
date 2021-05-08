package specifications

import (
	"errors"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
)

func NewIsCarModelCorrectSpecification(car *models.CarModel) error {
	isCarModelEmpty := car.Model == ""
	if isCarModelEmpty {
		return domain.NewError(
			errors.New("[domain][car][specifications][NewIsCarModelCorrectSpecification] car model is required"),
			domain.ValidationErrorCode,
			"Car model is required.",
		)
	}

	return nil
}
