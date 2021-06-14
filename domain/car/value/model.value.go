package value

import (
	"encoding/json"
	"fmt"
	"strings"

	"Sharykhin/rent-car/domain"
)

type (
	Model string
)

const (
	Audi Model = "audi"
	BMW        = "bmw"
)

// UnmarshalJSON implements Unmarshaler interface to convert json model representation into
// Model type
func (m *Model) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("[domain][car][types][Model][UnmarshalJSON] failed to unmarshal model type: %v", err)
	}
	model := Model(strings.ToLower(s))
	switch model {
	case Audi, BMW:
		*m = model
		return nil
	}

	return domain.NewError(
		fmt.Errorf("[domain][car][types][Model][UnmarshalJSON] car model is invalid: %v", m),
		domain.ValidationErrorCode,
		"Car model is invalid.",
	)
}
