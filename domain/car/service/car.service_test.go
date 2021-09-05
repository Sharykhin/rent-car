package service

import (
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car"
	"Sharykhin/rent-car/domain/car/factory"
	"Sharykhin/rent-car/domain/car/model"
	carMocks "Sharykhin/rent-car/domain/car/service/mocks"
	"Sharykhin/rent-car/domain/mocks"
)

func TestCarService_GetCarByID(t *testing.T) {
	Convey("Given a correct car ID", t, func() {
		carService, mockCarRepository, _, _ := newService()
		carID := domain.NewID()
		Convey("When a car with such ID exists", func() {
			mockCarRepository.On("GetCarByID", mock.Anything, carID).Return(&model.CarModel{
				ID: carID,
			}, nil)
			actualCar, actualErr := carService.GetCarByID(context.Background(), carID)
			Convey("Then the car is returned along with nil error", func() {
				So(actualErr, ShouldBeNil)
				So(actualCar.ID, ShouldEqual, carID)
			})
		})
	})
	Convey("Given a correct car ID", t, func() {

		carService, mockCarRepository, _, _ := newService()

		carID := domain.NewID()
		Convey("When a car with such ID does not exist", func() {

			mockCarRepository.On("GetCarByID", mock.Anything, carID).Return(nil, domain.NewError(car.ErrCarNotFound, "", domain.ResourceNotFoundErrorCode))
			actualCar, actualErr := carService.GetCarByID(context.TODO(), carID)

			Convey("Then a not found error is returned", func() {
				So(actualCar, ShouldBeNil)
				So(errors.Is(actualErr, car.ErrCarNotFound), ShouldBeTrue)
			})
		})
	})
}

func newService() (
	*CarService,
	*carMocks.CarRepositoryInterface,
	*mocks.TransactionInterface,
	*carMocks.FileStorageInterface,
) {
	mockCarRepository := new(carMocks.CarRepositoryInterface)
	mockTransactionSrv := new(mocks.TransactionInterface)
	mockFileStorage := new(carMocks.FileStorageInterface)

	engineFactory := new(factory.EngineValueFactory)
	carModelFactory := new(factory.CarModelFactory)

	return NewCarService(mockCarRepository, mockTransactionSrv, mockFileStorage, engineFactory, carModelFactory), mockCarRepository, mockTransactionSrv, mockFileStorage
}
