package models

import (
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car"
	consumerModels "Sharykhin/rent-car/domain/consumer/models"
)

type Requisition struct {
	ID        domain.ID
	Consumer  *consumerModels.Consumer
	Car       *car.CarModel
	DateFrom  time.Time
	DateTo    time.Time
	CreatedAt time.Time
}
