package domain

import (
	"fmt"
	"strings"
)

type (
	Code           string
	StackOperation string
	Error          struct {
		Err       error
		CallStack []StackOperation
		Code      Code
	}
)

const (
	ValidationErrorCode       Code = "VALIDATION"
	PayloadIsTooLarge              = "PAYLOAD_IS_TOO_LARGE"
	InternalServerErrorCode        = "INTERNAL_SERVER"
	ResourceNotFoundErrorCode      = "RESOURCE_NOT_FOUND"
)

func (e *Error) Error() string {
	ops := ""
	for i := len(e.CallStack) - 1; i >= 0; i-- {
		ops += string(e.CallStack[i]) + " "
	}
	ops = strings.TrimRight(ops, " ")

	return fmt.Sprintf("%s %s", ops, e.Err.Error())
}

func (e Error) Unwrap() error {
	return e.Err
}

func NewError(err error, operation StackOperation, code Code) *Error {
	return &Error{
		Err:       err,
		CallStack: []StackOperation{operation},
		Code:      code,
	}
}

// NewInternalError creates domain internal server error
func NewInternalError(err error, operation StackOperation) *Error {
	return NewError(err, operation, InternalServerErrorCode)
}

func WrapErrorWithStack(origin error, operation StackOperation) *Error {
	target, ok := origin.(*Error)
	if !ok {
		panic("origin error must be instance of Error")
	}
	target.CallStack = append(target.CallStack, operation)

	return target
}
