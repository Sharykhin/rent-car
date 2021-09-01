package service

import (
	"context"
	"errors"
	"testing"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car"
	"Sharykhin/rent-car/domain/car/factory"
	"Sharykhin/rent-car/domain/car/model"
	carMocks "Sharykhin/rent-car/domain/car/service/mocks"
	"Sharykhin/rent-car/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCarService_GetCarByID(t *testing.T) {
	// TODO: BDD doesn't look so nice with test table, try something new, like GoConvey
	tt := []struct {
		name        string
		carExists   bool
		inID        domain.ID
		expectedErr error
	}{
		{
			name: `
Given a correct car ID
When this car exists
Then the car is returned`,
			carExists:   true,
			inID:        domain.NewID(),
			expectedErr: nil,
		},
		{
			name: `
Give a correct car ID
When a car with such ID does not exist
Then a not found error is returned`,
			carExists:   false,
			inID:        domain.NewID(),
			expectedErr: car.ErrCarNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			srv, carRepository := newService()
			if tc.carExists {
				carRepository.On("GetCarByID", mock.Anything, mock.Anything).Return(&model.CarModel{
					ID: tc.inID,
				}, nil)
			} else {
				carRepository.On("GetCarByID", mock.Anything, mock.Anything).Return(nil, domain.NewError(car.ErrCarNotFound, "", domain.ResourceNotFoundErrorCode))
			}
			c, err := srv.GetCarByID(context.Background(), tc.inID)
			if tc.carExists {
				assert.Nil(t, err)
				assert.Equal(t, tc.inID, c.ID)
			} else {
				assert.True(t, errors.Is(err, car.ErrCarNotFound))
				assert.Nil(t, c)
			}

			carRepository.AssertExpectations(t)
		})
	}
}

func newService() (*CarService, *carMocks.CarRepositoryInterface) {
	carRepository := new(carMocks.CarRepositoryInterface)
	transactionSrv := new(mocks.TransactionInterface)
	fileStorage := new(carMocks.FileStorageInterface)

	engineFactory := new(factory.EngineValueFactory)
	carModelFactory := new(factory.CarModelFactory)

	return NewCarService(carRepository, transactionSrv, fileStorage, engineFactory, carModelFactory), carRepository
}
