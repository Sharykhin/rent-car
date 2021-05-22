package repositories

import (
	"context"
	"fmt"

	"github.com/lib/pq"

	"Sharykhin/rent-car/domain/requisition/models"
	"Sharykhin/rent-car/infrastructure/postgres"
)

type (
	// PostgresRequisitionRepository implements requisition repository
	PostgresRequisitionRepository struct {
		conn *postgres.Connection
	}
)

// NewPostgresRequisitionRepository creates a new instance of requisition repository
func NewPostgresRequisitionRepository(conn *postgres.Connection) *PostgresRequisitionRepository {
	repo := PostgresRequisitionRepository{
		conn: conn,
	}

	return &repo
}

// CreateRequisition creates a new requisition record
func (r *PostgresRequisitionRepository) CreateRequisition(
	ctx context.Context,
	requisition models.Requisition,
) (*models.Requisition, error) {
	_, err := r.conn.DB.ExecContext(
		ctx,
		"call rent_car($1, $2, $3, $4)",
		requisition.Car.ID,
		requisition.Consumer.ID,
		requisition.DateFrom,
		requisition.DateTo,
	)
	// TODO: Handle overlapping error
	if err != nil {
		pqErr := err.(*pq.Error)
		fmt.Println("pgErr", pqErr)
		return nil, fmt.Errorf("failed to call stored procedure rent_car, requisition %+v: %v", requisition, err)
	}

	return nil, err
}
