package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"Sharykhin/rent-car/domain"
)

type Model string

const (
	Audi Model = "Audi"
	BMW        = "BMW"
)

func (m *Model) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("[car][types][UnmarshalJSON] failed to unmarshal model type: %v", err)
	}
	model := Model(strings.Title(strings.ToLower(s)))
	switch model {
	case Audi, BMW:
		*m = model
		return nil
	}

	return domain.NewError(
		fmt.Errorf("[car][types][UnmarshalJSON] car model is invalid: %v", m),
		domain.ValidationErrorCode,
		"Car model is invalid.",
	)
}
