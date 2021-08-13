package factory

import (
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
	"Sharykhin/rent-car/domain/car/specification"
	"Sharykhin/rent-car/domain/car/value"
)

// NewCarModel creates a new car model with all validation steps
func NewCarModel(m value.Model, e *value.EngineValue) (*model.CarModel, error) {
	c := model.CarModel{
		ID:        domain.Empty(),
		Model:     m,
		Engine:    e,
		CreatedAt: time.Now().UTC(),
	}

	err := specification.IsCarModelCorrectSpecification(&c)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][factory][NewCarModel]")
	}

	return &c, nil
}
