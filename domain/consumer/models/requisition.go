package models

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
)

type Requisition struct {
	ID  domain.ID
	Car models.Car
}

func NewRequisition(car models.Car) Requisition {
	requisition := Requisition{
		ID:  domain.NewID(),
		Car: car,
	}

	return requisition
}
