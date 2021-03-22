package models

import (
	"Sharykhin/rent-car/domain"
	"time"
)

// Car represents a car that consumers will rent
type Car struct {
	ID        domain.ID
	Model     Model
	CreatedAt time.Time
}

// NewCar create a new car model
func NewCar(model Model) Car {
	car := Car{
		ID:        domain.ID(""),
		Model:     model,
		CreatedAt: time.Now().UTC(),
	}

	return car
}
