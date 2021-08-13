package service

import (
	"context"
	"encoding/json"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/dto"
	"Sharykhin/rent-car/domain/car/factory"
	"Sharykhin/rent-car/domain/car/model"
)

type (
	// CarService describes general business use-cases around car domain
	CarService struct {
		carRepo            CarRepositoryInterface
		transactionService domain.TransactionInterface
		fileStorage        FileStorageInterface
		engineValueFactory *factory.EngineValueFactory
	}
)

// NewCarService creates a new car service instance
func NewCarService(
	carRepo CarRepositoryInterface,
	transactionService domain.TransactionInterface,
	fileStorage FileStorageInterface,
	engineValueFactory *factory.EngineValueFactory,
) *CarService {
	srv := CarService{
		carRepo:            carRepo,
		transactionService: transactionService,
		fileStorage:        fileStorage,
		engineValueFactory: engineValueFactory,
	}

	return &srv
}

// CreateNewCar creates a new car in a database and then uploads it on s3
func (srv *CarService) CreateNewCar(ctx context.Context, dto *dto.CreateCarDto) (*model.CarModel, error) {
	engine, err := srv.engineValueFactory.CreateEngineValue(dto.Engine.Power, dto.Engine.IsTurbo)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][service][CarService][CreateNewCar]")
	}
	car, err := factory.NewCarModel(dto.Model, engine)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][service][CarService][CreateNewCar]")
	}

	car, err = srv.createCar(ctx, car)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][service][CarService][CreateNewCar]")
	}

	return car, nil
}

func (srv *CarService) createCar(ctx context.Context, car *model.CarModel) (*model.CarModel, error) {
	err := srv.transactionService.Wrap(ctx, func(ctx context.Context) error {
		var err error
		car, err = srv.carRepo.CreateCar(ctx, car)
		if err != nil {
			return domain.WrapErrorWithStack(err, "[domain][car][service][CarService][createCar]")
		}

		b, err := json.Marshal(&car)
		if err != nil {
			return domain.NewInternalError(
				fmt.Errorf("failed to marshal car: %v", err),
				"[domain][car][service][CarService][createCar]",
			)
		}

		err = srv.fileStorage.Upload(ctx, "cars/"+car.ID.String()+"/car.json", b)
		if err != nil {
			return domain.WrapErrorWithStack(err, "[domain][car][service][CarService][createCar]")
		}

		return nil
	})

	return car, err
}

// GetCarByID returns a specific car by its ID
func (srv *CarService) GetCarByID(ctx context.Context, ID domain.ID) (*model.CarModel, error) {
	c, err := srv.carRepo.GetCarByID(ctx, ID)

	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][car][service][CarService][GetCarByID]")
	}

	return c, err
}
