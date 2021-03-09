package models

import (
	"Sharykhin/rent-car/domain"
)

// Car represents a car that consumers will rent
type Car struct {
	ID    domain.ID
	Model Model
}

// NewCar create a new car model
func NewCar(model Model) Car {
	car := Car{
		ID:    domain.NewID(),
		Model: model,
	}

	return car
}
