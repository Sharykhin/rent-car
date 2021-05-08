package services

import (
	"context"
	"time"

	"Sharykhin/rent-car/domain"
	carModels "Sharykhin/rent-car/domain/car/models"
	consumerModels "Sharykhin/rent-car/domain/consumer/models"
	"Sharykhin/rent-car/domain/requisition/interfaces"
	"Sharykhin/rent-car/domain/requisition/models"
)

type (
	RequisitionService struct {
		requisitionRepo interfaces.RequisitionRepositoryInterface
	}
)

func NewRequisitionService(requisitionRep interfaces.RequisitionRepositoryInterface) *RequisitionService {
	srv := RequisitionService{
		requisitionRepo: requisitionRep,
	}

	return &srv
}

func (srv *RequisitionService) RentCar(ctx context.Context) (*models.Requisition, error) {
	requisition := models.Requisition{
		ID: domain.Empty(),
		Consumer: &consumerModels.ConsumerModel{
			ID:           domain.ID("8403116a-be3c-477d-a198-09f9adcda313"),
			FirstName:    "",
			LastName:     "",
			CreatedAt:    time.Now(),
			Requisitions: nil,
		},
		Car: &carModels.CarModel{
			ID: domain.ID("a87be964-770d-4af5-9269-b9874f1fadc0"),
		},
		DateFrom:  time.Now(),
		DateTo:    time.Now(),
		CreatedAt: time.Now(),
	}
	r, err := srv.requisitionRepo.CreateRequisition(ctx, requisition)

	return r, err
}