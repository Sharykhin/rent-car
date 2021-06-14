package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
	"Sharykhin/rent-car/infrastructure/postgres"
)

type (
	// PostgresCarRepository implements postgres car repository
	PostgresCarRepository struct {
		conn *postgres.Connection
	}
)

var (
	// ErrCarNotFound describes error when car was not found
	ErrCarNotFound = errors.New("car was not found")
)

// NewPostgresCarRepository creates a new car repository instance
func NewPostgresCarRepository(conn *postgres.Connection) *PostgresCarRepository {
	r := PostgresCarRepository{
		conn: conn,
	}

	return &r
}

// CreateCar creates a new car
func (r *PostgresCarRepository) CreateCar(ctx context.Context, car *model.CarModel) (*model.CarModel, error) {
	var id domain.ID
	stmt := `insert into public.cars(model, created_at) values($1, $2) returning id`
	err := r.conn.DB.QueryRowContext(ctx, stmt, car.Model, car.CreatedAt).Scan(&id)
	if err != nil {
		return nil, domain.NewInternalError(
			fmt.Errorf("failed to insert a new car record into cars table: %v", err),
			"[infrastructure][postgres][repositories][PostgresCarRepository][CreateCar]",
		)
	}

	newCar := model.CarModel{
		ID:        id,
		Model:     car.Model,
		CreatedAt: car.CreatedAt,
	}

	return &newCar, nil
}

// GetCarByID returns a car by its ID
func (r *PostgresCarRepository) GetCarByID(ctx context.Context, ID domain.ID) (*model.CarModel, error) {
	c := model.CarModel{}
	stmt := `select id, model, created_at from public.cars where id = $1`
	err := r.conn.DB.QueryRowContext(ctx, stmt, ID).Scan(&c.ID, &c.Model, &c.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewError(
				ErrCarNotFound,
				"[infrastructure][postgres][repositories][PostgresCarRepository][GetCarByID]",
				domain.ResourceNotFoundErrorCode,
			)
		}

		return nil, domain.NewInternalError(
			fmt.Errorf("failed to find a car in a database: %v", err),
			"[infrastructure][postgres][repositories][PostgresCarRepository][GetCarByID]",
		)
	}

	return &c, nil
}
