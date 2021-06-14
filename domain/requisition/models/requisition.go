package models

import (
	"time"

	"Sharykhin/rent-car/domain"
	carModels "Sharykhin/rent-car/domain/car/model"
	consumerModels "Sharykhin/rent-car/domain/consumer/models"
)

type Requisition struct {
	ID        domain.ID
	Consumer  *consumerModels.ConsumerModel
	Car       *carModels.CarModel
	DateFrom  time.Time
	DateTo    time.Time
	CreatedAt time.Time
}
