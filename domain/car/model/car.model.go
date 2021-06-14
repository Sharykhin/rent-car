package model

import (
	"encoding/json"
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/value"
)

type (
	// CarModel represents a car that consumers will rent
	CarModel struct {
		ID        domain.ID   `json:"id"`
		Model     value.Model `json:"model"`
		CreatedAt time.Time   `json:"created_at"`
	}
)

// MarshalJSON implements Marshaler interface to represent car model into json format
func (c *CarModel) MarshalJSON() ([]byte, error) {
	type Alias CarModel

	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(c),
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	})
}
