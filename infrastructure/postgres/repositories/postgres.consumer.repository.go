package repositories

import (
	"context"
	"fmt"

	"database/sql"
	"github.com/lib/pq"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
)

type (
	ConsumerRepository struct {
		db *sql.DB
	}
)

const (
	constraintUniqueCode = "23505"
)

// NewConsumerRepository is a function constructor that returns a new instance of the consumer repository
func NewConsumerRepository(db *sql.DB) *ConsumerRepository {
	r := ConsumerRepository{
		db: db,
	}

	return &r
}

func (r *ConsumerRepository) CreateConsumer(ctx context.Context, consumer *models.ConsumerModel) (*models.ConsumerModel, error) {
	var id domain.ID
	stmt := `insert into public.consumers(first_name, last_name, email, created_at) values($1, $2, $3, $4) returning id`

	err := r.db.QueryRowContext(ctx, stmt, consumer.FirstName, consumer.LastName, consumer.Email, consumer.CreatedAt).Scan(&id)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == constraintUniqueCode {
			return nil, domain.NewError(
				fmt.Errorf(
					"[infrastructure][postgres][repositories] failed to insert a new record into consumers table: %v",
					err,
				),
				domain.ValidationErrorCode,
				"Email is duplicated.",
			)
		}
		return nil, fmt.Errorf(
			"[infrastructure][postgres][repositories] failed to insert a new record into consumers table: %v",
			err,
		)
	}

	newConsumer := models.ConsumerModel{
		ID:           id,
		FirstName:    consumer.FirstName,
		LastName:     consumer.LastName,
		Email:        consumer.Email,
		CreatedAt:    consumer.CreatedAt,
		Requisitions: consumer.Requisitions,
	}
	consumer.ID = id

	return &newConsumer, nil
}
