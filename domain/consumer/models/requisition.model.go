package models

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
)

type Requisition struct {
	ID  domain.ID
	Car *model.CarModel
}

func NewRequisition(c *model.CarModel) *Requisition {
	requisition := Requisition{
		ID:  domain.Empty(),
		Car: c,
	}

	return &requisition
}
