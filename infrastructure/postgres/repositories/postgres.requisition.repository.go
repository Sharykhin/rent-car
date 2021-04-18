package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"Sharykhin/rent-car/domain/requisition/models"
)

type (
	PostgresRequisitionRepository struct {
		db *sql.DB
	}
)

// NewPostgresRequisitionRepository is a function constructor that returns a new instance of the consumer repository
func NewPostgresRequisitionRepository(db *sql.DB) *PostgresRequisitionRepository {
	repo := PostgresRequisitionRepository{
		db: db,
	}

	return &repo
}

func (repo *PostgresRequisitionRepository) CreateRequisition(ctx context.Context, requisition models.Requisition) (*models.Requisition, error) {
	_, err := repo.db.ExecContext(
		ctx,
		"call rent_car($1, $2, $3, $4)",
		requisition.Car.ID,
		requisition.Consumer.ID,
		requisition.DateFrom,
		requisition.DateTo,
	)
	if err != nil {
		pqErr := err.(*pq.Error)
		fmt.Println("pgErr", pqErr)
		return nil, fmt.Errorf("failed to call stored procedure rent_car, requisition %+v: %v", requisition, err)
	}

	return nil, err
}
