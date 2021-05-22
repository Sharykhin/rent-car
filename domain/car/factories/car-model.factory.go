package factories

import (
	"fmt"
	"time"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
	"Sharykhin/rent-car/domain/car/specifications"
	"Sharykhin/rent-car/domain/car/types"
)

// NewCarModel creates a new car model with all validation steps
func NewCarModel(m types.Model) (*models.CarModel, error) {
	c := models.CarModel{
		ID:        domain.Empty(),
		Model:     m,
		CreatedAt: time.Now().UTC(),
	}

	err := specifications.NewIsCarModelCorrectSpecification(&c)
	if err != nil {
		return nil, fmt.Errorf("[domain][factories][NewCarModel] failed to create a new car model: %w", err)
	}

	return &c, nil
}
