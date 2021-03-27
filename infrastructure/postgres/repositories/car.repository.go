package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
)

type CarRepository struct {
	db *sql.DB
}

// NewCarRepository is a function constructor that returns a new instance of car repository
func NewCarRepository(db *sql.DB) *CarRepository {
	r := CarRepository{
		db: db,
	}

	return &r
}

func (r *CarRepository) Create(ctx context.Context, car models.Car) (*models.Car, error) {
	var id domain.ID
	stmt := `insert into public.cars(model, created_at) values($1, $2) returning id`
	err := r.db.QueryRowContext(ctx, stmt, car.Model, car.CreatedAt).Scan(&id)
	if err != nil {
		return nil, domain.NewError(fmt.Errorf("failed to insert a new car record into cars table: %v", err), domain.InternalServerErrorCode, "Something went wrong.")
	}

	car.ID = id

	return &car, nil
}

func (r *CarRepository) GetCarByID(ctx context.Context, ID domain.ID) (*models.Car, error) {
	car := models.Car{}
	stmt := `select id, model, created_at from public.cars where id = $1`
	err := r.db.QueryRowContext(ctx, stmt, ID).Scan(&car.ID, &car.Model, &car.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewError(fmt.Errorf("failed to find a car in the database: %w", domain.ResourceNotFoundError), domain.ResourceNotFoundErrorCode, "Car resource was not found.")
		}

		return nil, domain.NewError(fmt.Errorf("failed to find a car in the database: %v", err), domain.InternalServerErrorCode, "Something went wrong.")

	}

	return &car, nil
}
