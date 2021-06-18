package model

import (
	"Sharykhin/rent-car/domain/requisition/value"
	"time"

	"Sharykhin/rent-car/domain"
	carModels "Sharykhin/rent-car/domain/car/model"
	consumerModels "Sharykhin/rent-car/domain/consumer/model"
)

type RequisitionModel struct {
	ID        domain.ID
	Consumer  *consumerModels.ConsumerModel
	Car       *carModels.CarModel
	Period    *value.Period
	CreatedAt time.Time
}
