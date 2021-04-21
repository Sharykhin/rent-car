package car

import (
	"encoding/json"
	"fmt"
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/specifications"
	"Sharykhin/rent-car/domain/car/types"
)

// CarModel represents a car that consumers will rent
type CarModel struct {
	ID        domain.ID   `json:"id"`
	Model     types.Model `json:"model"`
	CreatedAt time.Time   `json:"created_at"`
}

// TODO: this is a good way but is rather for view, between back to back it may be different
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

// NewCar create a new car model
func NewCarModel(m types.Model) (*CarModel, error) {
	c := CarModel{
		ID:        domain.Empty(),
		Model:     m,
		CreatedAt: time.Now().UTC(),
	}

	err := specifications.NewIsCarModelCorrectSpecification(&c)
	if err != nil {
		return nil, fmt.Errorf("[car][NewCarModel] failed to create a new car model: %w", err)
	}

	return &c, nil
}
