package response

import (
	"Sharykhin/rent-car/domain"
	"encoding/json"
	"log"
	"net/http"
)

type (
	WebSuccessResponse struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta"`
	}

	Error struct {
		Code    domain.Code `json:"code"`
		Message string      `json:"message"`
	}

	WebErrorResponse struct {
		Error Error `json:"error"`
	}
)

func Success(w http.ResponseWriter, data interface{}, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	r := WebSuccessResponse{
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
	r := WebSuccessResponse{
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

func BadRequest(w http.ResponseWriter, message string, code domain.Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	r := WebErrorResponse{
		Error: Error{
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

// TODO how to convert deep error context into client-friendly? For instance for 404 just return "resource was not found"
func Fail(w http.ResponseWriter, message string, code domain.Code) {
	w.Header().Set("Content-Type", "application/json")
	switch code {
	case domain.ResourceNotFoundErrorCode:
		w.WriteHeader(http.StatusNotFound)
	case domain.ValidationErrorCode:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	r := WebErrorResponse{
		Error: Error{
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

func NotFound(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	r := WebErrorResponse{
		Error: Error{
			Code:    domain.ResourceNotFoundErrorCode,
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

func Internal(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	r := WebErrorResponse{
		Error: Error{
			Code:    domain.InternalServerErrorCode,
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
