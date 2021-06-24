package domain

import "context"

type (
	// TransactionInterface describes api for wrapping execution flow into a transaction
	TransactionInterface interface {
		Wrap(ctx context.Context, fn func(context.Context) error) error
	}
)
