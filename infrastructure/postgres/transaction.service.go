package postgres

import (
	"context"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/logger"
)

type (
	// TransactionService provides an ability to wrap execution flow int transaction
	TransactionService struct {
		conn *Connection
	}
)

const (
	TXKey = "postgresTXKey"
)

// NewTransactionService creates a new instance of transaction service
func NewTransactionService(conn *Connection) *TransactionService {
	return &TransactionService{
		conn: conn,
	}
}

// Wrap wraps execution into transaction. It populates provided context with context value
// and put postgres open transaction instance into the context
func (t *TransactionService) Wrap(ctx context.Context, fn func(context.Context) error) error {
	logger.Log.Debugf("[infrastructure][postgres][TransactionService][Wrap] Begin transaction")
	tx, err := t.conn.DB.BeginTx(ctx, nil)
	if err != nil {
		return domain.NewInternalError(err, "[infrastructure][postgres][TransactionService][Wrap]")
	}

	txCtx := context.WithValue(ctx, TXKey, tx)
	err = fn(txCtx)
	if err != nil {
		logger.Log.Debugf("[infrastructure][postgres][TransactionService][Wrap] Rollback transaction")
		txErr := tx.Rollback()
		if txErr != nil {
			logger.Log.Errorf("[infrastructure][postgres][TransactionService][Wrap] failed to rollback a transaction: %v", txErr)
		}

		return err
	} else {
		logger.Log.Debugf("[infrastructure][postgres][TransactionService][Wrap] Commit transaction")
		txErr := tx.Commit()
		if txErr != nil {
			logger.Log.Errorf("[infrastructure][postgres][TransactionService][Wrap] failed to commit a transaction: %v", txErr)
		}

		return err
	}
}
