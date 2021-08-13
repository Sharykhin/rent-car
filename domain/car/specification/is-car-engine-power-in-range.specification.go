package specification

import (
	"errors"

	"Sharykhin/rent-car/domain"
)

var (
	ErrPowerIsTooHigh = errors.New("power can not be more than 100 when turbo is enabled")
)

const (
	maxPowerWhenTurboEnable = 100
)

func IsCarEnginePowerInRangeSpecification(power uint64, isTurbo bool) error {
	if isTurbo && power > maxPowerWhenTurboEnable {
		return domain.NewError(
			ErrPowerIsTooHigh,
			"[domain][car][specification][IsCarEnginePowerInRangeSpecification]",
			domain.ValidationErrorCode,
		)
	}

	return nil
}
