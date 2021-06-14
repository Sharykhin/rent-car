package query

import (
	"Sharykhin/rent-car/domain"
	"context"
	"fmt"

	"Sharykhin/rent-car/infrastructure/postgres"
)

type (
	PostgresCarQuery struct {
		conn *postgres.Connection
	}
	CarQueryEntity struct {
		ID string `json:"id"`
	}
)

func NewPostgresCarQuery(conn *postgres.Connection) *PostgresCarQuery {
	r := PostgresCarQuery{
		conn: conn,
	}

	return &r
}

func (r *PostgresCarQuery) GetPagedCarsList(ctx context.Context, limit, offset int) ([]CarQueryEntity, int, error) {
	rows, err := r.conn.DB.QueryContext(ctx, `select id from cars limit $1 offset $2`, limit, offset)

	errChan := make(chan error)
	defer close(errChan)
	totalChan := make(chan int)
	defer close(totalChan)

	go func(ctx context.Context, totalChan chan<- int, errChan chan<- error) {
		var total int
		err := r.conn.DB.QueryRowContext(ctx, "select count(*) from cars").Scan(&total)
		if err != nil {
			errChan <- domain.NewInternalError(
				fmt.Errorf("failed to make a count query: %v", err),
				"[infrastructure][postgres][query][PostgresCarQuery][GetPagedCarsList]",
			)
			return
		}
		totalChan <- total
	}(ctx, totalChan, errChan)

	if err != nil {
		return nil, 0, domain.NewInternalError(
			fmt.Errorf("failed to make a select to get cars list: %v", err),
			"[infrastructure][postgres][query][PostgresCarQuery][GetPagedCarsList]",
		)
	}

	defer rows.Close()

	cars := make([]CarQueryEntity, 0)

	for rows.Next() {
		var car CarQueryEntity
		err := rows.Scan(&car.ID)
		if err != nil {
			return nil, 0, domain.NewInternalError(
				fmt.Errorf("failed to scan car: %v", err),
				"[infrastructure][postgres][query][PostgresCarQuery][GetPagedCarsList]",
			)
		}
		cars = append(cars, car)
	}

	var total int
	select {
	case total = <-totalChan:
	case err := <-errChan:
		return nil, 0, err
	}

	return cars, total, nil
}
