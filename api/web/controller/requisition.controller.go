package controller

import (
	"Sharykhin/rent-car/domain/requisition/services"
	"net/http"

	"Sharykhin/rent-car/api/web/response"
)

type (
	RequisitionController struct {
		requisitionSrv *services.RequisitionService
	}
)

func NewRequisitionController(requisitionSrv *services.RequisitionService) *RequisitionController {
	ctrl := RequisitionController{
		requisitionSrv: requisitionSrv,
	}

	return &ctrl
}

func (ctrl *RequisitionController) CreateRequisition(w http.ResponseWriter, r *http.Request) {
	req, err := ctrl.requisitionSrv.RentCar(r.Context())
	if err != nil {
		response.Fail(w, err)
		return
	}

	response.Created(w, req, nil)
}
