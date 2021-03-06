package specification

import (
	"errors"
	"net"
	"regexp"
	"strings"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/model"
)

var (
	ErrConsumerEmailRequired = errors.New("email is required")
	ErrConsumerEmailInvalid  = errors.New("email is invalid")

	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// IsConsumerEmailCorrectSpecification validates consumer email
func IsConsumerEmailCorrectSpecification(consumer *model.ConsumerModel) error {
	isEmpty := consumer.Email == ""
	if isEmpty {
		return domain.NewError(
			ErrConsumerEmailRequired,
			"[domain][consumer][specification][IsConsumerEmailCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	if !isEmailValid(consumer.Email) {
		return domain.NewError(
			ErrConsumerEmailInvalid,
			"[domain][consumer][specification][IsConsumerEmailCorrectSpecification]",
			domain.ValidationErrorCode,
		)
	}

	return nil
}

// isEmailValid checks if the email provided passes the required structure
// and length test. It also checks the domain has a valid MX record.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}

	if !emailRegex.MatchString(e) {
		return false
	}

	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])

	if err != nil || len(mx) == 0 {
		return false
	}

	return true
}
