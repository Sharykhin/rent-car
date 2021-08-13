package factory

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/specification"
	"Sharykhin/rent-car/domain/car/value"
)

type (
	EngineValueFactory struct {
		isDebug bool
	}
)

func NewEngineValueFactory(isDebug bool) *EngineValueFactory {
	f := EngineValueFactory{
		isDebug: isDebug,
	}

	return &f
}

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
