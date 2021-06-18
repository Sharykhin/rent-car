package model

import (
	"time"

	"Sharykhin/rent-car/domain"
)

// ConsumerModel is a person who rents a car
type ConsumerModel struct {
	ID        domain.ID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}
