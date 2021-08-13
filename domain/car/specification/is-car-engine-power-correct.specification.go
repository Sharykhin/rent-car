package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
)

var (
	ErrPowerIsZero = errors.New("power can not be zero")
)

func IsCarEnginePowerCorrectSpecification(power uint64) error {
	if power == 0 {
		return domain.NewError(
			ErrPowerIsZero,
			"[domain][car][specification][IsCarEnginePowerInRangeSpecification]",
			domain.ValidationErrorCode,
		)
	}

	return nil
}
