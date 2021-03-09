package models

import (
	"Sharykhin/rent-car/domain"
	"encoding/json"
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
		return err
	}
	model := Model(strings.Title(strings.ToLower(s)))
	switch model {
	case Audi, BMW:
		*m = model
		return nil
	}
	return domain.InvalidCarModelError
}
