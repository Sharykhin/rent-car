package value

import (
	"Sharykhin/rent-car/domain"
	"errors"
)

type (
	Period struct {
		StartAt domain.Date `json:"start_at"`
		EndAt   domain.Date `json:"end_at"`
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
