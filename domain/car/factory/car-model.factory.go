package factory

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
	"Sharykhin/rent-car/domain/car/value"
)

type (
	// CarModelFactory is responsible for creating a car domain model
	CarModelFactory struct {
	}
)

// NewCarModelFactory creates a new instance of car model factory
func NewCarModelFactory() *CarModelFactory {
	f := CarModelFactory{}

	return &f
}

// CreateCar creates a new car
func (c *CarModelFactory) CreateCar(carModel value.Model, engine *value.EngineValue) (*model.CarModel, error) {
	car, err := model.NewCarModel(domain.Empty(), carModel, engine)
	if err != nil {
		return car, domain.WrapErrorWithStack(err, "[domain][car][factory][CarModelFactory][CreateCar]")
	}

	return car, err
}
