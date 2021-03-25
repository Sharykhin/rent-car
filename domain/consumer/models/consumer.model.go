package models

import (
	"Sharykhin/rent-car/domain"
	carModels "Sharykhin/rent-car/domain/car/models"
)

// Consumer is a person/organization who/what rents cars
type Consumer struct {
	ID           domain.ID
	Name         string
	Requisitions []Requisition
}

func (c *Consumer) RentCar(car carModels.Car) error {
	if len(c.Requisitions) > 2 {
		return domain.RequisitionLimitExceededError
	}

	requisition := NewRequisition(car)
	c.Requisitions = append(c.Requisitions, *requisition)

	return nil
}

// NewConsumer creates a new consumer
func NewConsumer(name string, requisitions []Requisition) *Consumer {
	consumer := Consumer{
		ID:           domain.Empty(),
		Name:         name,
		Requisitions: requisitions,
	}

	return &consumer
}
