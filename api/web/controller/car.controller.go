package controller

import (
	"net/http"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/api/web/util"
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/dto"
	"Sharykhin/rent-car/domain/car/service"
	"Sharykhin/rent-car/domain/car/value"
)

type (
	// CarController is a web controller that handles API requests around car domain model
	CarController struct {
		carService *service.CarService
	}

	// CreateCarPayload this is a request body for creating a new car
	CreateCarPayload struct {
		Model  value.Model `json:"model"`
		Engine struct {
			Power   uint64 `json:"power"`
			IsTurbo bool   `json:"is_turbo"`
		} `json:"engine"`
	}
)

func NewCarController(carSrv *service.CarService) *CarController {
	ctrl := CarController{
		carService: carSrv,
	}

	return &ctrl
}

// CreateCar handles http request to create a new car
func (ctrl *CarController) CreateCar(w http.ResponseWriter, r *http.Request) {
	var payload CreateCarPayload
	err := util.DecodeJSONBody(w, r, &payload)
	if err != nil {
		response.Fail(w, domain.WrapErrorWithStack(err, "[api][web][controller][CarController][CreateCar]"))
		return
	}

	car, err := ctrl.carService.CreateNewCar(r.Context(), &dto.CreateCarDto{
		Model: payload.Model,
		Engine: dto.EngineDto{
			Power:   payload.Engine.Power,
			IsTurbo: payload.Engine.IsTurbo,
		},
	})
	if err != nil {
		response.Fail(w, domain.WrapErrorWithStack(err, "[api][web][controller][CarController][CreateCar]"))
		return
	}

	response.Created(w, car, nil)
}

// GetCarByID handles rest api to get a car by its id
func (ctrl *CarController) GetCarByID(w http.ResponseWriter, r *http.Request) {
	ID := getUrlParam(r, "id")

	c, err := ctrl.carService.GetCarByID(r.Context(), domain.ID(ID))

	if err != nil {
		response.Fail(w, domain.WrapErrorWithStack(err, "[api][web][controller][CarController][GetCarByID]"))
		return
	}

	response.Success(w, c, nil)

}
