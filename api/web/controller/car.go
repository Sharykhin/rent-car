package controller

import (
	"Sharykhin/rent-car/domain"
	"encoding/json"
	"errors"
	"net/http"

	"Sharykhin/rent-car/domain/car/models"
)

type (
	CarController struct {
	}

	CreateCarPayload struct {
		Model models.Model `json:"model"`
	}

	CreateResponse struct {
		ID string `json:"id"`
	}

	FailResponse struct {
		Message string `json:"message"`
	}
)

func (c *CarController) CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var payload CreateCarPayload
	err := decoder.Decode(&payload)
	if err != nil {
		if errors.Is(err, domain.InvalidCarModelError) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		_ = json.NewEncoder(w).Encode(FailResponse{
			Message: err.Error(),
		})
		return
	}

	car := models.NewCar(payload.Model)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(CreateResponse{
		ID: car.ID.String(),
	})
}
