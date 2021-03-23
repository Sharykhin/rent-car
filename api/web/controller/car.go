package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
	"Sharykhin/rent-car/domain/car/services"
)

type (
	CarController struct {
		carService *services.CarService
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

func NewCarController(carService *services.CarService) *CarController {
	ctrl := CarController{
		carService: carService,
	}

	return &ctrl
}

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

	car, err := c.carService.CreateNewCar(r.Context(), payload.Model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(FailResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&car)
}
