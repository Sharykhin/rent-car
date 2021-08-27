package car

import "errors"

var (
	// ErrCarNotFound describes error when car was not found
	ErrCarNotFound = errors.New("car was not found")
)
