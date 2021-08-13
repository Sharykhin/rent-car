package value

import (
	"encoding/json"

	"Sharykhin/rent-car/domain"
)

type (
	EngineValue struct {
		Power   uint64 `json:"power"`
		IsTurbo bool   `json:"is_turbo"`
	}
)

func (e *EngineValue) UnmarshalJSON(b []byte) error {
	var payload EngineValue
	err := json.Unmarshal(b, &payload)
	if err != nil {
		return domain.NewInternalError(err, "[domain][car][value][EngineValue][UnmarshalJSON]")
	}
	e.Power = 100
	e.IsTurbo = true

	return nil
}
