package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"Sharykhin/rent-car/api/web/response"
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
			response.BadRequest(w, err.Error(), err.Code)
			return
		}
		response.Internal(w, err.Error())
		return
	}

	car, err := c.carService.CreateNewCar(r.Context(), payload.Model)
	if err != nil {
		response.Internal(w, err.Error())
	}

	response.Created(w, car, nil)
}

func (c *CarController) GetCarByID(w http.ResponseWriter, r *http.Request) {
	ID := getUrlParam(r, "id")

	car, err := c.carService.GetCarByID(r.Context(), domain.ID(ID))

	if err != nil {
		if err, ok := err.(*domain.Error); ok {
			response.Fail(w, err.Error(), err.Code)
			return
		}
		response.Internal(w, err.Error())
		return
	}

	response.Success(w, car, nil)

}

func getUrlParam(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}
