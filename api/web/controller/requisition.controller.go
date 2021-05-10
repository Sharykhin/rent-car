package controller

import (
	"fmt"
	"net/http"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/api/web/util"
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/requisition/services"
)

type (
	RequisitionController struct {
		requisitionSrv *services.RequisitionService
	}

	CreateRequisitionPayload struct {
		CarID      domain.ID   `json:"car_id"`
		ConsumerID domain.ID   `json:"consumer_id"`
		StartAt    domain.Date `json:"start_at"`
		EndAt      domain.Date `json:"end_at"`
	}
)

func NewRequisitionController(requisitionSrv *services.RequisitionService) *RequisitionController {
	ctrl := RequisitionController{
		requisitionSrv: requisitionSrv,
	}

	return &ctrl
}

func (ctrl *RequisitionController) CreateRequisition(w http.ResponseWriter, r *http.Request) {
	var payload CreateRequisitionPayload
	err := util.DecodeJSONBody(w, r, &payload)
	fmt.Println(payload)
	if err != nil {
		response.Fail(w, err)
		return
	}

	//req, err := ctrl.requisitionSrv.RentCar2(r.Context())
	//if err != nil {
	//	response.Fail(w, err)
	//	return
	//}

	response.Created(w, nil, nil)
}
