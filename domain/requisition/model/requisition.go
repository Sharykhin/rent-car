package model

import (
	"time"

	"Sharykhin/rent-car/domain"
	carModels "Sharykhin/rent-car/domain/car/model"
	consumerModels "Sharykhin/rent-car/domain/consumer/model"
	"Sharykhin/rent-car/domain/requisition/value"
)

// RequisitionModel is an order of renting a car
type RequisitionModel struct {
	ID        domain.ID                     `json:"id"`
	Consumer  *consumerModels.ConsumerModel `json:"consumer"`
	Car       *carModels.CarModel           `json:"car"`
	Period    *value.Period                 `json:"period"`
	CreatedAt time.Time                     `json:"created_at"`
}
