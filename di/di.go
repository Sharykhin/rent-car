package di

import (
	"errors"
	"fmt"
	"os"

	"Sharykhin/rent-car/api/web/controller"
	carServices "Sharykhin/rent-car/domain/car/services"
	consumerServices "Sharykhin/rent-car/domain/consumer/services"
	requisitionServices "Sharykhin/rent-car/domain/requisition/services"
	"Sharykhin/rent-car/infrastructure/postgres"
	postgresRepos "Sharykhin/rent-car/infrastructure/postgres/repositories"
	log "github.com/sirupsen/logrus"
)

var (
	initialized             = false
	AlreadyInitializedError = errors.New("[di] di has already been initialized")
	Container               *container
)

type (
	// Container is a service locator that keeps all the instances of the application
	container struct {
		PostgresConn                  *postgres.Connection
		PostgresCarRepository         *postgresRepos.PostgresCarRepository
		PostgresConsumerRepository    *postgresRepos.PostgresConsumerRepository
		PostgresRequisitionRepository *postgresRepos.PostgresRequisitionRepository
		CarService                    *carServices.CarService
		ConsumerService               *consumerServices.ConsumerService
		RequisitionService            *requisitionServices.RequisitionService
		CarController                 *controller.CarController
		ConsumerController            *controller.ConsumerController
		RequisitionController         *controller.RequisitionController
		Logger                        *log.Entry
	}
)

// Init initializes a new container. It acts as function constructor but Init sounds here more semantical
// cause you are able to initialize it only once.
func Init() error {
	if initialized {
		return AlreadyInitializedError
	}

	postgresConn, err := postgres.NewConnection(os.Getenv("POSTGRES_URL"))
	if err != nil {
		return fmt.Errorf("[di] failed to create a new postgers connection instance: %v", err)
	}

	postgresCarRepository := postgresRepos.NewPostgresCarRepository(postgresConn)
	postgresConsumerRepository := postgresRepos.NewPostgresConsumerRepository(postgresConn)
	postgresRequisitionRepository := postgresRepos.NewPostgresRequisitionRepository(postgresConn)

	carService := carServices.NewCarService(postgresCarRepository)
	consumerService := consumerServices.NewConsumerService(postgresConsumerRepository)
	requisitionService := requisitionServices.NewRequisitionService(postgresRequisitionRepository)

	carController := controller.NewCarController(carService)
	consumerController := controller.NewConsumerController(consumerService)
	requisitionController := controller.NewRequisitionController(requisitionService)

	// TODO: @improve set json formatter and rename msg to message and time to timestamp
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatalf("[web][main] failed to parse a log level env variable: %v", err)
	}
	log.SetLevel(level)
	logger := log.WithField("service", os.Getenv("SERVICE_ID"))

	container := container{
		PostgresConn:                  postgresConn,
		PostgresCarRepository:         postgresCarRepository,
		PostgresConsumerRepository:    postgresConsumerRepository,
		PostgresRequisitionRepository: postgresRequisitionRepository,
		CarService:                    carService,
		ConsumerService:               consumerService,
		RequisitionService:            requisitionService,
		CarController:                 carController,
		ConsumerController:            consumerController,
		RequisitionController:         requisitionController,
		Logger:                        logger,
	}

	initialized = true

	Container = &container

	return nil
}
