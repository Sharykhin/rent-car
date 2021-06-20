package controller

import (
	"net/http"

	"Sharykhin/rent-car/api/web/response"
	"Sharykhin/rent-car/api/web/util"
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/requisition/service"
)

type (
	// RequisitionController handles incoming http request
	// and describes API around requisition
	RequisitionController struct {
		requisitionSrv *service.RequisitionService
	}

	// CreateRequisitionPayload is a request body for creating a new requisition
	CreateRequisitionPayload struct {
		CarID      domain.ID   `json:"car_id"`
		ConsumerID domain.ID   `json:"consumer_id"`
		StartAt    domain.Date `json:"start_at"`
		EndAt      domain.Date `json:"end_at"`
	}
)

// NewRequisitionController is a function constructor that creates a new instance of requisition controller
func NewRequisitionController(requisitionSrv *service.RequisitionService) *RequisitionController {
	ctrl := RequisitionController{
		requisitionSrv: requisitionSrv,
	}

	return &ctrl
}

// CreateRequisition handles endpoint of creating a new requisition
func (ctrl *RequisitionController) CreateRequisition(w http.ResponseWriter, r *http.Request) {
	var payload CreateRequisitionPayload
	err := util.DecodeJSONBody(w, r, &payload)
	if err != nil {
		response.Fail(w, domain.WrapErrorWithStack(err, "[api][web][controller][RequisitionController][CreateRequisition]"))
		return
	}

	requisition, err := ctrl.requisitionSrv.RentCar(r.Context(), payload.CarID, payload.ConsumerID, payload.StartAt, payload.EndAt)
	if err != nil {
		response.Fail(w, domain.WrapErrorWithStack(err, "[api][web][controller][RequisitionController][CreateRequisition]"))
		return
	}

	response.Created(w, requisition, nil)
}
