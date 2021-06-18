package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"Sharykhin/rent-car/domain"
)

const (
	MaxPayloadInBytes = 1048576
)

// DecodeJSONBody parses request body into destination struct
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	r.Body = http.MaxBytesReader(w, r.Body, MaxPayloadInBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return domain.NewError(
				fmt.Errorf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset),
				"[api][web][util][DecodeJSONBody]",
				domain.ValidationErrorCode,
			)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return domain.NewError(
				errors.New("request body contains badly-formed JSON"),
				"[api][web][util][DecodeJSONBody]",
				domain.ValidationErrorCode,
			)

		case errors.As(err, &unmarshalTypeError):
			return domain.NewError(
				fmt.Errorf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset),
				"[api][web][util][DecodeJSONBody]",
				domain.ValidationErrorCode,
			)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return domain.NewError(
				fmt.Errorf("request body contains unknown field %s", fieldName),
				"[api][web][util][DecodeJSONBody]",
				domain.ValidationErrorCode,
			)

		case errors.Is(err, io.EOF):
			return domain.NewError(
				errors.New("request body must not be empty"),
				"[api][web][util][DecodeJSONBody]",
				domain.ValidationErrorCode,
			)
		case err.Error() == "http: request body too large":
			return domain.NewError(
				errors.New("request body must not be larger than 1MB"),
				"[api][web][util][DecodeJSONBody]",
				domain.PayloadIsTooLarge,
			)
		default:
			return domain.NewInternalError(err, "[api][web][util][DecodeJSONBody]")
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return domain.NewError(
			errors.New("request body must only contain a single JSON object"),
			"[api][web][util][DecodeJSONBody]",
			domain.ValidationErrorCode,
		)
	}

	return nil
}
