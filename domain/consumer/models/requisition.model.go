package models

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
)

type Requisition struct {
	ID  domain.ID
	Car *models.CarModel
}

func NewRequisition(c *models.CarModel) *Requisition {
	requisition := Requisition{
		ID:  domain.Empty(),
		Car: c,
	}

	return &requisition
}
