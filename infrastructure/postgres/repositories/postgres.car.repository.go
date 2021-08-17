package repositories

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/car/model"
	"Sharykhin/rent-car/domain/car/value"
	"Sharykhin/rent-car/infrastructure/postgres"
)

type (
	// PostgresCarRepository implements postgres car repository
	PostgresCarRepository struct {
		conn *postgres.Connection
	}

	carProperties struct {
		Engine carEngine `json:"engine"`
	}
	carEngine struct {
		Power   uint64 `json:"power"`
		IsTurbo bool   `json:"is_turbo"`
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

// Value makes the carProperties struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a carProperties) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the carProperties struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *carProperties) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return domain.NewInternalError(
			errors.New("type assertion to []byte failed"),
			"[infrastructure][postgres][repositories][carProperties][Scan]",
		)
	}

	return json.Unmarshal(b, &a)
}

// CreateCar creates a new car
func (r *PostgresCarRepository) CreateCar(ctx context.Context, car *model.CarModel) (*model.CarModel, error) {
	var id domain.ID
	var err error

	stmt := `insert into public.cars(model, properties, created_at) values($1, $2, $3) returning id`
	props := carProperties{
		Engine: carEngine{
			Power:   car.Engine.Power,
			IsTurbo: car.Engine.IsTurbo,
		},
	}
	tx, ok := ctx.Value(postgres.TXKey).(*sql.Tx)
	if ok {
		err = tx.QueryRowContext(ctx, stmt, car.Model, props, car.CreatedAt).Scan(&id)
	} else {
		err = r.conn.DB.QueryRowContext(ctx, stmt, car.Model, props, car.CreatedAt).Scan(&id)
	}

	if err != nil {
		return nil, domain.NewInternalError(
			fmt.Errorf("failed to insert a new car record into cars table: %v", err),
			"[infrastructure][postgres][repositories][PostgresCarRepository][CreateCar]",
		)
	}

	newCar := model.CarModel{
		ID:        id,
		Model:     car.Model,
		Engine:    car.Engine,
		CreatedAt: car.CreatedAt,
	}

	return &newCar, nil
}

// GetCarByID returns a car by its ID
func (r *PostgresCarRepository) GetCarByID(ctx context.Context, ID domain.ID) (*model.CarModel, error) {
	car := model.CarModel{}
	props := carProperties{}
	stmt := `select id, model, properties, created_at from public.cars where id = $1`
	err := r.conn.DB.QueryRowContext(ctx, stmt, ID).Scan(&car.ID, &car.Model, &props, &car.CreatedAt)

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

	car.Engine = &value.EngineValue{
		Power:   props.Engine.Power,
		IsTurbo: props.Engine.IsTurbo,
	}

	return &car, nil
}
