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

func (r *CarRepository) CreateCar(ctx context.Context, car *models.CarModel) (*models.CarModel, error) {
	var id domain.ID
	stmt := `insert into public.cars(model, created_at) values($1, $2) returning id`
	err := r.db.QueryRowContext(ctx, stmt, car.Model, car.CreatedAt).Scan(&id)
	if err != nil {
		return nil, domain.NewInternalError(
			fmt.Errorf("[infrastructure][postgres][repositories] failed to insert a new car record into cars table: %v", err),
		)
	}

	nc := models.CarModel{
		ID:        id,
		Model:     car.Model,
		CreatedAt: car.CreatedAt,
	}

	return &nc, nil
}

func (r *CarRepository) GetCarByID(ctx context.Context, ID domain.ID) (*models.CarModel, error) {
	c := models.CarModel{}
	stmt := `select id, model, created_at from public.cars where id = $1`
	err := r.db.QueryRowContext(ctx, stmt, ID).Scan(&c.ID, &c.Model, &c.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewError(
				fmt.Errorf("[infrastructure][postgres][repositories] failed to find a car in the database: %v", err),
				domain.ResourceNotFoundErrorCode,
				"Car resource was not found.",
			)
		}

		return nil, domain.NewInternalError(
			fmt.Errorf("[infrastructure][postgres][repositories] failed to insert a new car record into cars table: %v", err),
		)
	}

	return &c, nil
}
