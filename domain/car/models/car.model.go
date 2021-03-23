package models

import (
	"encoding/json"
	"time"

	"Sharykhin/rent-car/domain"
)

// Car represents a car that consumers will rent
type Car struct {
	ID        domain.ID `json:"id"`
	Model     Model     `json:"model"`
	CreatedAt time.Time `json:"created_at"`
}

// TODO: this is a good way but is rather for view between back to back it may be different
func (c *Car) MarshalJSON() ([]byte, error) {
	type Alias Car

	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(c),
		CreatedAt: c.CreatedAt.Format("Mon Jan _2"),
	})
}

// NewCar create a new car model
func NewCar(model Model) *Car {
	car := Car{
		ID:        domain.ID(""),
		Model:     model,
		CreatedAt: time.Now().UTC(),
	}

	return &car
}
