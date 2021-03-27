package controller

import (
	"encoding/json"
	"net/http"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
	"Sharykhin/rent-car/domain/car/services"
)

type (
	// CarController is a web controller that handles API requests around car domain model
	CarController struct {
		carService *services.CarService
	}

	// CreateCarPayload this is a request body for creating a new car
	CreateCarPayload struct {
		Model models.Model `json:"model"`
	}
)

func NewCarController(carService *services.CarService) *CarController {
	ctrl := CarController{
		carService: carService,
	}

	return &ctrl
}

func (c *CarController) CreateCar(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload CreateCarPayload
	err := decoder.Decode(&payload)

	if err != nil {
		if err, ok := err.(*domain.Error); ok {
			response.Fail(w, err.Message, err.Code)
			return
		}
		response.Internal(w, err.Error())
		return
	}

	car, err := c.carService.CreateNewCar(r.Context(), payload.Model)
	if err != nil {
		if err, ok := err.(*domain.Error); ok {
			response.Fail(w, err.Message, err.Code)
			return
		}
		response.Internal(w, err.Error())
		return
	}

	response.Created(w, car, nil)
}

func (c *CarController) GetCarByID(w http.ResponseWriter, r *http.Request) {
	ID := getUrlParam(r, "id")

	car, err := c.carService.GetCarByID(r.Context(), domain.ID(ID))

	if err != nil {
		if err, ok := err.(*domain.Error); ok {
			response.Fail(w, err.Message, err.Code)
			return
		}
		response.Internal(w, err.Error())
		return
	}

	response.Success(w, car, nil)

}
