package repositories

import (
	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type ConsumerRepository struct {
	db *sql.DB
}

// NewConsumerRepository is a function constructor that returns a new instance of the consumer repository
func NewConsumerRepository(db *sql.DB) *ConsumerRepository {
	r := ConsumerRepository{
		db: db,
	}

	return &r
}

func (r *ConsumerRepository) Create(ctx context.Context, consumer models.Consumer) (*models.Consumer, error) {
	var id domain.ID
	stmt := `insert into public.consumers(first_name, last_name, email, created_at) values($1, $2, $3, $4) returning id`

	err := r.db.QueryRowContext(ctx, stmt, consumer.FirstName, consumer.LastName, consumer.Email, consumer.CreatedAt).Scan(&id)
	if err != nil {
		pqErr := err.(*pq.Error)
		fmt.Println(pqErr.Code == "23505", pqErr.Code)
		if pqErr.Code == "23505" {
			return nil, domain.NewError(fmt.Errorf("failed to insert a new record into consumers table: email is duplicated: %v", err), domain.ValidationErrorCode)
		}
		return nil, domain.NewError(fmt.Errorf("failed to insert a new record into consumers table: %v", err), domain.InternalServerErrorCode)
	}

	consumer.ID = id

	return &consumer, nil
}
