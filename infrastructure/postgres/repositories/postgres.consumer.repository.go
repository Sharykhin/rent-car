package repositories

import (
	"context"
	"fmt"

	"github.com/lib/pq"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
	"Sharykhin/rent-car/infrastructure/postgres"
)

type (
	// PostgresConsumerRepository implement consumer repository
	PostgresConsumerRepository struct {
		conn *postgres.Connection
	}
)

const (
	constraintUniqueCode = "23505"
)

// NewPostgresConsumerRepository creates a new consumer repository instance
func NewPostgresConsumerRepository(conn *postgres.Connection) *PostgresConsumerRepository {
	r := PostgresConsumerRepository{
		conn: conn,
	}

	return &r
}

// CreateConsumer creates a new consumer
func (r *PostgresConsumerRepository) CreateConsumer(
	ctx context.Context,
	consumer *models.ConsumerModel,
) (*models.ConsumerModel, error) {
	var id domain.ID
	stmt := `insert into public.consumers(first_name, last_name, email, created_at) values($1, $2, $3, $4) returning id`

	err := r.conn.DB.QueryRowContext(ctx, stmt, consumer.FirstName, consumer.LastName, consumer.Email, consumer.CreatedAt).Scan(&id)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == constraintUniqueCode {
			return nil, domain.NewError(
				fmt.Errorf(
					"[PostgresConsumerRepository][CreateConsumer] failed to insert a new record into consumers table: %v",
					err,
				),
				domain.ValidationErrorCode,
				"Email is duplicated.",
			)
		}
		return nil, fmt.Errorf("[PostgresConsumerRepository][CreateConsumer] failed to insert a new record into consumers table: %v", err)
	}

	newConsumer := models.ConsumerModel{
		ID:           id,
		FirstName:    consumer.FirstName,
		LastName:     consumer.LastName,
		Email:        consumer.Email,
		CreatedAt:    consumer.CreatedAt,
		Requisitions: consumer.Requisitions,
	}

	return &newConsumer, nil
}
