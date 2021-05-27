package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
	"Sharykhin/rent-car/infrastructure/postgres"
)

type (
	// PostgresCarRepository implements postgres car repository
	PostgresCarRepository struct {
		conn *postgres.Connection
	}
)

// NewPostgresCarRepository creates a new car repository instance
func NewPostgresCarRepository(conn *postgres.Connection) *PostgresCarRepository {
	r := PostgresCarRepository{
		conn: conn,
	}

	return &r
}

// CreateCar creates a new car
func (r *PostgresCarRepository) CreateCar(ctx context.Context, car *models.CarModel) (*models.CarModel, error) {
	var id domain.ID
	stmt := `insert into public.cars(model, created_at) values($1, $2) returning id`
	err := r.conn.DB.QueryRowContext(ctx, stmt, car.Model, car.CreatedAt).Scan(&id)
	if err != nil {
		return nil, domain.NewInternalError(
			fmt.Errorf("[PostgresCarRepository][CreateCar] failed to insert a new car record into cars table: %v", err),
		)
	}

	newCar := models.CarModel{
		ID:        id,
		Model:     car.Model,
		CreatedAt: car.CreatedAt,
	}

	return &newCar, nil
}

// GetCarByID returns a car model by its ID
func (r *PostgresCarRepository) GetCarByID(ctx context.Context, ID domain.ID) (*models.CarModel, error) {
	c := models.CarModel{}
	stmt := `select id, model, created_at from public.cars where id = $1`
	err := r.conn.DB.QueryRowContext(ctx, stmt, ID).Scan(&c.ID, &c.Model, &c.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewError(
				fmt.Errorf("[PostgresCarRepository][GetCarByID] failed to find a car in the database: %v", err),
				domain.ResourceNotFoundErrorCode,
				"Car resource was not found.",
			)
		}

		return nil, fmt.Errorf("[PostgresCarRepository][GetCarByID] failed to find a car in the database: %v", err)
	}

	return &c, nil
}
