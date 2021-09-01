package specification

import (
	"errors"
	"testing"

	"Sharykhin/rent-car/domain/car/value"
	"github.com/stretchr/testify/assert"
)

func TestIsCarModelCorrectSpecification(t *testing.T) {
	tt := []struct {
		name        string
		inModel     value.Model
		expectedErr error
	}{
		{
			name: `
Given an empty string as car model
When car model is required
Then an error is returned`,
			inModel:     "",
			expectedErr: ErrCarModelRequired,
		},
		{
			name: `
Given an unexpected car model
When it is not not support one
Then an error is returned`,
			inModel:     value.Model("unexpected"),
			expectedErr: ErrCarModelInvalid,
		},
		{
			name: `
Given a valid car model
When it is supported one
When nil error is returned`,
			inModel:     value.BMW,
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := IsCarModelCorrectSpecification(tc.inModel)
			if err != nil {
				assert.True(t, errors.Is(err, tc.expectedErr))
			} else {
				assert.Equal(t, tc.expectedErr, err)
			}
		})
	}
}
