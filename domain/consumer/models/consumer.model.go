package models

import (
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
)

// ConsumerModel is a person/organization who/what rents cars
type ConsumerModel struct {
	ID           domain.ID
	FirstName    string
	LastName     string
	Email        string
	CreatedAt    time.Time
	Requisitions []Requisition
}

func (cs *ConsumerModel) RentCar(c *models.CarModel) error {
	if len(cs.Requisitions) > 2 {
		return domain.RequisitionLimitExceededError
	}

	requisition := NewRequisition(c)
	cs.Requisitions = append(cs.Requisitions, *requisition)

	return nil
}
