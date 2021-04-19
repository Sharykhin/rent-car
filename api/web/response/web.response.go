package response

import (
	"Sharykhin/rent-car/api/web/util"
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

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
	var malformedErr *util.MalformedRequest
	if errors.As(err, &domainErr) {
		asDomainError(w, domainErr.Message, domainErr.Code, err)
		return
	}
	if errors.As(err, &malformedErr) {
		asNativeError(w, malformedErr.Status, malformedErr.Msg, domain.ValidationErrorCode, err)
		return
	}
	asDomainError(w, err.Error(), domain.InternalServerErrorCode, err)
}

func asNativeError(w http.ResponseWriter, status int, msg string, code domain.Code, origin error) {
	switch status {
	case http.StatusInternalServerError:
		log.Error(origin)
	}

	w.WriteHeader(status)
	r := errorResponse{
		Error: webError{
			Code:    code,
			Message: msg,
		},
	}
	sendErrorResponse(w, &r)
}

func asDomainError(w http.ResponseWriter, message string, code domain.Code, origin error) {
	switch code {
	case domain.ResourceNotFoundErrorCode:
		w.WriteHeader(http.StatusNotFound)
	case domain.ValidationErrorCode:
		w.WriteHeader(http.StatusBadRequest)
	default:
		log.Error(origin)
		w.WriteHeader(http.StatusInternalServerError)
	}

	r := errorResponse{
		Error: webError{
			Code:    code,
			Message: message,
		},
	}
	sendErrorResponse(w, &r)
}

func sendErrorResponse(w http.ResponseWriter, r *errorResponse) {
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
