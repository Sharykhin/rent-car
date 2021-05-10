package domain

import (
	"fmt"
	"strings"
	"time"
)

type (
	// Date is a custom datetime type with supports several formats
	Date time.Time
)

var formats = []string{time.RFC3339, "2006-01-02"}

// UnmarshalJSON implements unmarshaler interface to part provided date into one of supported formats
// otherwise returns an error
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	for _, format := range formats {
		date, err := time.Parse(format, s)
		if err != nil {
			continue
		}

		*d = Date(date)
		return nil
	}

	return NewError(
		fmt.Errorf("provided date %s is has invalid format", s),
		ValidationErrorCode,
		"Date format is invalid",
	)
}

// String implements stringer interface to represent custom Date type into string
func (d Date) String() string {
	return time.Time(d).Format("2006-01-02 15:04:05")
}
