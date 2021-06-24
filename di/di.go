package di

import (
	"errors"
	"fmt"
	"os"

	"Sharykhin/rent-car/api/web/controller"
	carService "Sharykhin/rent-car/domain/car/service"
	consumerServices "Sharykhin/rent-car/domain/consumer/service"
	requisitionService "Sharykhin/rent-car/domain/requisition/service"
	"Sharykhin/rent-car/infrastructure/postgres"
	postgresRepos "Sharykhin/rent-car/infrastructure/postgres/repositories"
	"Sharykhin/rent-car/infrastructure/s3"
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
		CarService                    *carService.CarService
		ConsumerService               *consumerServices.ConsumerService
		RequisitionService            *requisitionService.RequisitionService
		CarController                 *controller.CarController
		ConsumerController            *controller.ConsumerController
		RequisitionController         *controller.RequisitionController
		Logger                        *log.Entry
		PostgresTransactionService    *postgres.TransactionService
		S3Client                      *s3.Client
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
	postgresTransactionService := postgres.NewTransactionService(postgresConn)
	postgresCarRepository := postgresRepos.NewPostgresCarRepository(postgresConn)
	postgresConsumerRepository := postgresRepos.NewPostgresConsumerRepository(postgresConn)
	postgresRequisitionRepository := postgresRepos.NewPostgresRequisitionRepository(postgresConn)

	isS3ForcePathStyle := os.Getenv("AWS_S3_FORCE_PATH_STYLE") == "true"
	s3Client := s3.NewClient(os.Getenv("AWS_S3_ENDPOINT"), isS3ForcePathStyle, os.Getenv("AWS_S3_BUCKET_NAME"))

	carSrv := carService.NewCarService(postgresCarRepository, postgresTransactionService, s3Client)
	consumerService := consumerServices.NewConsumerService(postgresConsumerRepository, postgresTransactionService, s3Client)
	requisitionSrv := requisitionService.NewRequisitionService(
		postgresRequisitionRepository,
		postgresCarRepository,
		postgresConsumerRepository,
	)

	carController := controller.NewCarController(carSrv)
	consumerController := controller.NewConsumerController(consumerService)
	requisitionController := controller.NewRequisitionController(requisitionSrv)

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
		CarService:                    carSrv,
		ConsumerService:               consumerService,
		RequisitionService:            requisitionSrv,
		CarController:                 carController,
		ConsumerController:            consumerController,
		RequisitionController:         requisitionController,
		Logger:                        logger,
		PostgresTransactionService:    postgresTransactionService,
		S3Client:                      s3Client,
	}

	initialized = true

	Container = &container

	return nil
}
