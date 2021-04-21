package models

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car"
)

type Requisition struct {
	ID  domain.ID
	Car *car.CarModel
}

func NewRequisition(c *car.CarModel) *Requisition {
	requisition := Requisition{
		ID:  domain.Empty(),
		Car: c,
	}

	return &requisition
}
