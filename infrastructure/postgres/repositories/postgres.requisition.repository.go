package repositories

import (
	"context"
	"errors"
	"fmt"

	"strings"

	"github.com/lib/pq"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/requisition/model"
	"Sharykhin/rent-car/infrastructure/postgres"
)

type (
	// PostgresRequisitionRepository implements requisition repository
	PostgresRequisitionRepository struct {
		conn *postgres.Connection
	}
)

var (
	ErrPeriodOverlapping = errors.New("provided period is overlapping")
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
	requisition *model.RequisitionModel,
) (*model.RequisitionModel, error) {
	var newID domain.ID
	err := r.conn.DB.QueryRowContext(
		ctx,
		"call rent_car($1, $2, $3, $4, $5)",
		nil,
		requisition.Car.ID,
		requisition.Consumer.ID,
		requisition.Period.StartAt.String(),
		requisition.Period.EndAt.String(),
	).Scan(&newID)

	if err != nil {
		pqErr := err.(*pq.Error)
		if strings.Contains(pqErr.Message, "ERR_OVERLAPPING") {
			return nil, domain.NewError(
				ErrPeriodOverlapping,
				"[infrastructure][postgres][repositories][PostgresRequisitionRepository][CreateRequisition]",
				domain.ValidationErrorCode,
			)
		}

		return nil, domain.NewInternalError(
			fmt.Errorf("failed to call stored procedure rent_car, requisition %+v: %v", requisition, err),
			"[infrastructure][postgres][repositories][PostgresRequisitionRepository][CreateRequisition]",
		)
	}

	newRequisition := model.RequisitionModel{
		ID:        newID,
		Car:       requisition.Car,
		Consumer:  requisition.Consumer,
		Period:    requisition.Period,
		CreatedAt: requisition.CreatedAt,
	}

	return &newRequisition, err
}
