package models

import (
	"time"

	"Sharykhin/rent-car/domain"
	carModels "Sharykhin/rent-car/domain/car/models"
	consumerModels "Sharykhin/rent-car/domain/consumer/models"
)

type Requisition struct {
	ID        domain.ID
	Consumer  *consumerModels.Consumer
	Car       *carModels.Car
	DateFrom  time.Time
	DateTo    time.Time
	CreatedAt time.Time
}
