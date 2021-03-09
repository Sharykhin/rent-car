package domain

import "errors"

var (
	RequisitionLimitExceededError = errors.New("requisition limit exceeded")
	InvalidCarModelError          = errors.New("car model is invalid")
)
