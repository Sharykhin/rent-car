package controller

import (
	"net/http"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/api/web/util"
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/service"
	"Sharykhin/rent-car/domain/car/types"
)

type (
	// CarController is a web controller that handles API requests around car domain model
	CarController struct {
		carService *service.CarService
	}

	// CreateCarPayload this is a request body for creating a new car
	CreateCarPayload struct {
		Model types.Model `json:"model"`
	}
)

func NewCarController(carSrv *service.CarService) *CarController {
	ctrl := CarController{
		carService: carSrv,
	}

	return &ctrl
}

func (ctrl *CarController) CreateCar(w http.ResponseWriter, r *http.Request) {
	var payload CreateCarPayload
	err := util.DecodeJSONBody(w, r, &payload)

	if err != nil {
		response.Fail(w, err)
		return
	}

	c, err := ctrl.carService.CreateNewCar(r.Context(), payload.Model)
	if err != nil {
		response.Fail(w, err)
		return
	}

	response.Created(w, c, nil)
}

func (ctrl *CarController) GetCarByID(w http.ResponseWriter, r *http.Request) {
	ID := getUrlParam(r, "id")

	c, err := ctrl.carService.GetCarByID(r.Context(), domain.ID(ID))

	if err != nil {
		response.Fail(w, err)
		return
	}

	response.Success(w, c, nil)

}
