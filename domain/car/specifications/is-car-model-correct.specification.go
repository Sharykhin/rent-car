package specifications

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
	"errors"
)

func NewIsCarModelCorrectSpecification(car *models.Car) error {
	if car.Model == "" {
		return domain.NewError(errors.New("car model is required"), domain.ValidationErrorCode, "Car model is required.")
	}

	return nil
}
