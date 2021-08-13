package dto

import "Sharykhin/rent-car/domain/car/value"

type (
	// CreateCarDto describes all the necessary data for creating a new car
	CreateCarDto struct {
		Model  value.Model
		Engine EngineDto
	}
	// EngineDto describes all the necessary data to create engine value object
	EngineDto struct {
		Power   uint64
		IsTurbo bool
	}
)
