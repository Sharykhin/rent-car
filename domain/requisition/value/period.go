package value

import (
	"Sharykhin/rent-car/domain"
	"errors"
)

type (
	Period struct {
		StartAt domain.Date
		EndAt   domain.Date
	}
)

var (
	ErrPeriodIsEmpty = errors.New("provided period has no time range")
)

func NewPeriod(startAt, endAt domain.Date) (*Period, error) {
	if startAt == endAt {
		return nil, domain.NewError(
			ErrPeriodIsEmpty,
			"[domain][requisition][value][NewPeriod]",
			domain.ValidationErrorCode,
		)
	}

	p := Period{
		StartAt: startAt,
		EndAt:   endAt,
	}

	return &p, nil
}
