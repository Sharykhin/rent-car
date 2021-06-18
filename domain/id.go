package domain

import (
	"errors"
	"strings"

	guuid "github.com/google/uuid"
)

type (
	// ID is a general Unique Identifier that is used across the whole application for each entity
	ID string
)

var (
	ErrIDInvalid = errors.New("id is not valid")
)

func (id *ID) UnmarshalJSON(b []byte) error {
	s, err := ParseID(string(b))
	if err != nil {
		return WrapErrorWithStack(err, "[domain][ID][UnmarshalJSON]")
	}

	*id = s

	return nil
}

// String returns string representation of ID
func (id ID) String() string {
	return string(id)
}

func (id ID) IsEmpty() bool {
	return strings.Trim(id.String(), " ") == ""
}

// ParseID validates an id and return custom ID if there are no errors
func ParseID(id string) (ID, error) {
	guid, err := guuid.Parse(id)
	if err != nil {
		return "", NewError(ErrIDInvalid, "[domain][ParseID]", ValidationErrorCode)
	}

	return ID(guid.String()), nil
}

// NewID generate a new ID. It uses guid as unique identifier
func NewID() ID {
	guid := guuid.New()

	return ID(guid.String())
}

// TODO: @improve rename to NewEmptyID cause just Empty under domain package is not clear what it means
func Empty() ID {
	return ID("")
}
