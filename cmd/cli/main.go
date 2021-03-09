package main

import (
	"fmt"

	carModels "Sharykhin/rent-car/domain/car/models"
	consumerModels "Sharykhin/rent-car/domain/consumer/models"
)

func main() {
	car := carModels.NewCar(carModels.Audi)
	consumer := consumerModels.NewConsumer("John", make([]consumerModels.Requisition, 0))
	requisition := consumerModels.NewRequisition(car)
	fmt.Println(car, consumer, requisition.Car)
}
