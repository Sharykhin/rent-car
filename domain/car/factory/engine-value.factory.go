package factory

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/specification"
	"Sharykhin/rent-car/domain/car/value"
)

type (
	// EngineValueFactory is responsible for creating car engine value object
	// isDebug flag does nothing but shows how we can inject any potential dependencies
	EngineValueFactory struct {
		isDebug bool
	}
)

// NewEngineValueFactory creates a new instance of engine value object factory
func NewEngineValueFactory(isDebug bool) *EngineValueFactory {
	f := EngineValueFactory{
		isDebug: isDebug,
	}

	return &f
}

// CreateEngineValue creates a new car engine value object
func (f *EngineValueFactory) CreateEngineValue(power uint64, isTurbo bool) (*value.EngineValue, error) {
	e := value.EngineValue{
		Power:   power,
		IsTurbo: isTurbo,
	}

	err := specification.IsCarEnginePowerCorrectSpecification(e.Power)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][factory][EngineValueFactory][CreateEngineValue]")
	}

	err = specification.IsCarEnginePowerInRangeSpecification(e.Power, e.IsTurbo)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][factory][EngineValueFactory][CreateEngineValue]")
	}

	return &e, nil
}
