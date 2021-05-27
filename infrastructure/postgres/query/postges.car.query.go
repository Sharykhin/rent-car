package query

import (
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

	go func(totalChan chan<- int, errChan chan<- error) {
		var total int
		err := r.conn.DB.QueryRowContext(ctx, "select count(*) from cars").Scan(&total)
		if err != nil {
			errChan <- err
		}
		totalChan <- total
	}(totalChan, errChan)

	if err != nil {
		return nil, 0, fmt.Errorf("[PostgresCarQuery][GetPagedCarsList] failed to make a select to get cars list: %v", err)
	}

	defer rows.Close()

	cars := make([]CarQueryEntity, 0)

	for rows.Next() {
		var car CarQueryEntity
		err := rows.Scan(&car.ID)
		if err != nil {
			return nil, 0, fmt.Errorf("[PostgresCarQuery][GetPagedCarsList] failed to scan car: %v", err)
		}
		cars = append(cars, car)
	}

	var total int
	select {
	case total = <-totalChan:
	case err := <-errChan:
		return nil, 0, fmt.Errorf("faield: %v", err)
	}

	return cars, total, nil
}
