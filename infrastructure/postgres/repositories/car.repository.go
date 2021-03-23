package repositories

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/models"
	"context"
	"database/sql"
	"fmt"
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
		return nil, fmt.Errorf("failed to insert a new record into cars table: %v", err)
	}

	car.ID = id

	return &car, nil
}
