package model

import (
	"time"

	"Sharykhin/rent-car/domain"
)

// ConsumerModel is a person who rents a car
type ConsumerModel struct {
	ID        domain.ID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
