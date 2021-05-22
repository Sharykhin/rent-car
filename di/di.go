package di

import (
	"errors"
	"fmt"
	"os"

	"Sharykhin/rent-car/infrastructure/postgres"
	postgresRepos "Sharykhin/rent-car/infrastructure/postgres/repositories"
)

var (
	initialized             = false
	AlreadyInitializedError = errors.New("[di] di has already been initialized")
)

type (
	// Container is a service locator that keeps all the instances of the application
	Container struct {
		PostgresConn                  *postgres.Connection
		PostgresCarRepository         *postgresRepos.PostgresCarRepository
		PostgresConsumerRepository    *postgresRepos.PostgresConsumerRepository
		PostgresRequisitionRepository *postgresRepos.PostgresRequisitionRepository
	}
)

// Init initializes a new container. It acts as function constructor but Init sounds here more semantical
// cause you are able to initialize it only once.
func Init() (*Container, error) {
	if initialized {
		return nil, AlreadyInitializedError
	}

	postgresConn, err := postgres.NewConnection(os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, fmt.Errorf("[di] failed to create a new postgers connection instance: %v", err)
	}

	postgresCarRepository := postgresRepos.NewPostgresCarRepository(postgresConn)
	postgresConsumerRepository := postgresRepos.NewPostgresConsumerRepository(postgresConn)
	postgresRequisitionRepository := postgresRepos.NewPostgresRequisitionRepository(postgresConn)

	container := Container{
		PostgresConn:                  postgresConn,
		PostgresCarRepository:         postgresCarRepository,
		PostgresConsumerRepository:    postgresConsumerRepository,
		PostgresRequisitionRepository: postgresRequisitionRepository,
	}

	initialized = true

	return &container, nil
}
