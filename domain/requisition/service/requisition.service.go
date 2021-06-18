package service

import (
	"Sharykhin/rent-car/domain/requisition/value"
	"context"
	"time"

	"Sharykhin/rent-car/domain"
	carServices "Sharykhin/rent-car/domain/car/service"
	consumerServices "Sharykhin/rent-car/domain/consumer/service"
	"Sharykhin/rent-car/domain/requisition/model"
)

type (
	RequisitionService struct {
		requisitionRepo RequisitionRepositoryInterface
		carRepo         carServices.CarRepositoryInterface
		consumerRepo    consumerServices.ConsumerRepositoryInterface
	}
)

func NewRequisitionService(
	requisitionRepo RequisitionRepositoryInterface,
	carRepo carServices.CarRepositoryInterface,
	consumerRepo consumerServices.ConsumerRepositoryInterface,
) *RequisitionService {
	srv := RequisitionService{
		requisitionRepo: requisitionRepo,
		carRepo:         carRepo,
		consumerRepo:    consumerRepo,
	}

	return &srv
}

func (srv *RequisitionService) RentCar(
	ctx context.Context,
	carID domain.ID,
	consumerID domain.ID,
	startAt domain.Date,
	endAt domain.Date,
) (*model.RequisitionModel, error) {
	car, err := srv.carRepo.GetCarByID(ctx, carID)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][requisition][service][RequisitionService][RentCar]")
	}

	consumer, err := srv.consumerRepo.GetConsumerByID(ctx, consumerID)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][requisition][service][RequisitionService][RentCar]")
	}
	period, err := value.NewPeriod(startAt, endAt)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][requisition][service][RequisitionService][RentCar]")
	}
	requisition := model.RequisitionModel{
		ID:        domain.Empty(),
		Car:       car,
		Consumer:  consumer,
		Period:    period,
		CreatedAt: time.Now(),
	}

	r, err := srv.requisitionRepo.CreateRequisition(ctx, &requisition)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][requisition][service][RequisitionService][RentCar]")
	}

	return r, nil
}
