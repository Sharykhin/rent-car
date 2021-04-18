package interfaces

import (
	"context"

	"Sharykhin/rent-car/domain/requisition/models"
)

type (
	RequisitionRepositoryInterface interface {
		CreateRequisition(ctx context.Context, requisition models.Requisition) (*models.Requisition, error)
	}
)
