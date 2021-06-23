package domain

import "context"

type (
	TransactionInterface interface {
		Wrap(ctx context.Context, fn func(context.Context) error) error
	}
)
