package model

import (
	"encoding/json"
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/specification"
	"Sharykhin/rent-car/domain/car/value"
)

type (
	// CarModel represents a car that consumers will rent
	CarModel struct {
		ID        domain.ID          `json:"id"`
		Model     value.Model        `json:"model"`
		Engine    *value.EngineValue `json:"engine"`
		CreatedAt time.Time          `json:"created_at"`
	}
)

// NewCarModel creates a new car model with all validation rules applied
// TODO: @concern the way how to pass ID looks pretty weird like it exists but empty reckon if there are better solutions
func NewCarModel(ID domain.ID, model value.Model, engine *value.EngineValue) (*CarModel, error) {
	c := CarModel{
		ID:        ID,
		Model:     model,
		Engine:    engine,
		CreatedAt: time.Now().UTC(),
	}

	if err := specification.IsCarModelCorrectSpecification(c.Model); err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][model][NewCarModel]")
	}

	return &c, nil
}

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
