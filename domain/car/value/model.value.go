package value

import (
	"encoding/json"
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
		return domain.NewInternalError(err, "[domain][car][types][Model][UnmarshalJSON]")
	}
	model := Model(strings.ToLower(s))
	*m = model

	return nil
}
