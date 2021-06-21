package value

import (
	"errors"

	"Sharykhin/rent-car/domain"
)

type (
	// Period represents general period value object
	Period struct {
		StartAt domain.Date `json:"start_at"`
		EndAt   domain.Date `json:"end_at"`
	}
)

var (
	ErrPeriodIsEmpty     = errors.New("provided period has no time range")
	ErrStartDateAfterEnd = errors.New("start date is after end date")
)

// NewPeriod creates a new period including general validation rules
func NewPeriod(startAt, endAt domain.Date) (*Period, error) {
	if startAt == endAt {
		return nil, domain.NewError(
			ErrPeriodIsEmpty,
			"[domain][requisition][value][NewPeriod]",
			domain.ValidationErrorCode,
		)
	}

	if startAt.After(endAt) {
		return nil, domain.NewError(
			ErrStartDateAfterEnd,
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
