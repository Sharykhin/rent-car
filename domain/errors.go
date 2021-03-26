package domain

import (
	"errors"
	"fmt"
)

type (
	Code  string
	Error struct {
		Code Code
		Err  error
	}
)

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}

const (
	ValidationErrorCode       Code = "VALIDATION"
	InternalServerErrorCode        = "INTERNAL_SERVER"
	ResourceNotFoundErrorCode      = "RESOURCE_NOT_FOUND"
)

var (
	ResourceNotFoundError         = errors.New("resource was not found")
	RequisitionLimitExceededError = errors.New("requisition limit exceeded")
	InvalidCarModelError          = errors.New("car model is invalid")
)

func NewError(err error, code Code) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

func WrapError(err error, target error) *Error {
	targetErr, ok := target.(*Error)
	if !ok {
		panic(fmt.Errorf("provided error is not domain.Error type but: %T. Origin error: %v", err, err))
	}

	targetErr.Err = fmt.Errorf("%v: %w", err, targetErr.Err)
	return targetErr
}
