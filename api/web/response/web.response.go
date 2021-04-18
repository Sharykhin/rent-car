package response

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"Sharykhin/rent-car/domain"
)

type (
	successResponse struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta"`
	}

	webError struct {
		Code    domain.Code `json:"code"`
		Message string      `json:"message"`
	}
	errorResponse struct {
		Error webError `json:"error"`
	}
)

func Success(w http.ResponseWriter, data interface{}, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	r := successResponse{
		Data: data,
		Meta: meta,
	}
	err := json.NewEncoder(w).Encode(&r)
	if err != nil {
		log.Printf("failed to encode http response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Internal Server Error"))
		if err != nil {
			log.Printf("failed to write http response: %v", err)
		}
	}

}

func Created(w http.ResponseWriter, data interface{}, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	r := successResponse{
		Data: data,
		Meta: meta,
	}
	err := json.NewEncoder(w).Encode(&r)
	if err != nil {
		log.Printf("failed to encode http response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Internal Server Error"))
		if err != nil {
			log.Printf("failed to write http response: %v", err)
		}
	}

}

func Fail(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	var domainErr *domain.Error
	if errors.As(err, &domainErr) {
		asDomainError(w, domainErr.Message, domainErr.Code)
		return
	}
	asDomainError(w, err.Error(), domain.InternalServerErrorCode)
}

func asDomainError(w http.ResponseWriter, message string, code domain.Code) {
	w.Header().Set("Content-Type", "application/json")
	switch code {
	case domain.ResourceNotFoundErrorCode:
		w.WriteHeader(http.StatusNotFound)
	case domain.ValidationErrorCode:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	r := errorResponse{
		Error: webError{
			Code:    code,
			Message: message,
		},
	}
	err := json.NewEncoder(w).Encode(&r)
	if err != nil {
		log.Printf("failed to encode http response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Internal Server Error"))
		if err != nil {
			log.Printf("failed to write http response: %v", err)
		}
	}
}
