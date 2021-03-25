package response

import (
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
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	WebErrorResponse struct {
		Error Error `json:"error"`
	}
)

func Created(w http.ResponseWriter, data interface{}, meta interface{}) {
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

func BadRequest(w http.ResponseWriter, message string, code string) {
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

func Internal(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	r := WebErrorResponse{
		Error: Error{
			Code:    "INTERNAL_SERVER_ERROR",
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
