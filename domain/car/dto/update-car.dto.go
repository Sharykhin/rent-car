package dto

import "Sharykhin/rent-car/domain/car/value"

type (
	// UpdateCarDto describes data that are used to update the car
	UpdateCarDto struct {
		Model  value.Model
		Engine EngineValueDto
	}
)
