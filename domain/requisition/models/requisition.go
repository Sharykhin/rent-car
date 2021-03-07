package models

import (
	"time"

	carModels "Sharykhin/rent-car/domain/car/models"
	consumerModels "Sharykhin/rent-car/domain/consumer/models"
)

type Requisition struct {
	ID       string
	Consumer consumerModels.Consumer
	Car      carModels.Car
	ExpireAt time.Time
}
