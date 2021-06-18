package service

import (
	"context"

	"Sharykhin/rent-car/domain/requisition/model"
)

type (
	RequisitionRepositoryInterface interface {
		CreateRequisition(ctx context.Context, requisition *model.RequisitionModel) (*model.RequisitionModel, error)
	}
)
