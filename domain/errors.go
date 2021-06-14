package domain

import (
	"fmt"
)

type (
	Code  string
	Error struct {
		Code    Code
		Err     error
		Message string
	}
)

func (e Error) Error() string {
	return fmt.Sprintf("code - %v, error - %s ", e.Code, e.Err.Error())
}

func (e Error) Unwrap() error {
	return e.Err
}

const (
	ValidationErrorCode       Code = "VALIDATION"
	InternalServerErrorCode        = "INTERNAL_SERVER"
	ResourceNotFoundErrorCode      = "RESOURCE_NOT_FOUND"
)

// NewError creates a new domain error with corresponding code and semantical message
func NewError(err error, code Code, message string) *Error {
	return &Error{
		Code:    code,
		Err:     err,
		Message: message,
	}
}

// NewInternalError creates domain internal server error
func NewInternalError(err error) *Error {
	return NewError(err, InternalServerErrorCode, "Something went wrong")
}
