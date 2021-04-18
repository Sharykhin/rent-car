package models

import (
	"Sharykhin/rent-car/domain"
	"encoding/json"
	"fmt"
	"strings"
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
		return fmt.Errorf("failed to unmarshal model type")
	}
	model := Model(strings.Title(strings.ToLower(s)))
	switch model {
	case Audi, BMW:
		*m = model
		return nil
	}

	return domain.NewError(domain.InvalidCarModelError, domain.ValidationErrorCode, "Car model is invalid.")
}
