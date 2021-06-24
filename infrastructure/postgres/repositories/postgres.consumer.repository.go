package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/model"
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

var (
	ErrEmailAlreadyExists = errors.New("email is already in use")
	// ErrConsumerNotFound describes error when car was not found
	ErrConsumerNotFound = errors.New("consumer was not found")
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
	consumer *model.ConsumerModel,
) (*model.ConsumerModel, error) {
	var id domain.ID
	var err error

	stmt := `insert into public.consumers(first_name, last_name, email, created_at) values($1, $2, $3, $4) returning id`

	txVal := ctx.Value(postgres.TXKey)
	tx, ok := txVal.(*sql.Tx)
	if ok {
		err = tx.QueryRowContext(ctx, stmt, consumer.FirstName, consumer.LastName, consumer.Email, consumer.CreatedAt).Scan(&id)
	} else {
		err = r.conn.DB.QueryRowContext(ctx, stmt, consumer.FirstName, consumer.LastName, consumer.Email, consumer.CreatedAt).Scan(&id)
	}

	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == constraintUniqueCode {
			return nil, domain.NewError(
				ErrEmailAlreadyExists,
				"[infrastructure][postgres][repositories][PostgresConsumerRepository][CreateConsumer]",
				domain.ValidationErrorCode,
			)
		}

		return nil, domain.NewInternalError(
			fmt.Errorf("failed to insert a new record into consumers table: %v", err),
			"[infrastructure][postgres][repositories][PostgresConsumerRepository][CreateConsumer]",
		)
	}

	newConsumer := model.ConsumerModel{
		ID:        id,
		FirstName: consumer.FirstName,
		LastName:  consumer.LastName,
		Email:     consumer.Email,
		CreatedAt: consumer.CreatedAt,
	}

	return &newConsumer, nil
}

func (r *PostgresConsumerRepository) GetConsumerByID(ctx context.Context, ID domain.ID) (*model.ConsumerModel, error) {
	consumer := model.ConsumerModel{}
	stmt := `select id, first_name, last_name, email, created_at from public.consumers where id = $1`
	err := r.conn.DB.QueryRowContext(ctx, stmt, ID).Scan(
		&consumer.ID,
		&consumer.FirstName,
		&consumer.LastName,
		&consumer.Email,
		&consumer.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewError(
				ErrConsumerNotFound,
				"[infrastructure][postgres][repositories][PostgresConsumerRepository][GetConsumerByID]",
				domain.ResourceNotFoundErrorCode,
			)
		}

		return nil, domain.NewInternalError(
			fmt.Errorf("failed to find a consumer in a database: %v", err),
			"[infrastructure][postgres][repositories][PostgresConsumerRepository][GetConsumerByID]",
		)
	}

	return &consumer, nil
}
