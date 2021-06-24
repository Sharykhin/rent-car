package service

import (
	"context"
	"encoding/json"
	"fmt"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/domain/consumer/factory"
	"Sharykhin/rent-car/domain/consumer/model"
)

type (
	// ConsumerService describes business use-cases around consumer domain
	ConsumerService struct {
		consumerRepo       ConsumerRepositoryInterface
		transactionService domain.TransactionInterface
		fileStorage        FileStorageInterface
	}
)

// NewConsumerService is a function constructor that creates a new instance of ConsumerService struct
func NewConsumerService(
	consumerRepo ConsumerRepositoryInterface,
	transactionService domain.TransactionInterface,
	fileStorage FileStorageInterface,
) *ConsumerService {
	srv := ConsumerService{
		consumerRepo:       consumerRepo,
		transactionService: transactionService,
		fileStorage:        fileStorage,
	}

	return &srv
}

// CreateNewConsumer creates a new consumer
func (srv *ConsumerService) CreateNewConsumer(ctx context.Context, firstName, lastName, email string) (*model.ConsumerModel, error) {
	consumer, err := factories.NewConsumerModel(firstName, lastName, email)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][consumer][service][ConsumerService][CreateNewConsumer][NewConsumerModel]")
	}

	consumer, err = srv.createConsumer(ctx, consumer)
	if err != nil {
		return nil, domain.WrapErrorWithStack(err, "[domain][consumer][service][ConsumerService][CreateNewConsumer][CreateConsumer]")
	}

	return consumer, nil
}

func (srv *ConsumerService) createConsumer(ctx context.Context, consumer *model.ConsumerModel) (*model.ConsumerModel, error) {
	err := srv.transactionService.Wrap(ctx, func(subCtx context.Context) error {
		var err error
		consumer, err = srv.consumerRepo.CreateConsumer(subCtx, consumer)
		if err != nil {
			return domain.WrapErrorWithStack(err, "[domain][consumer][service][ConsumerService][createConsumer]")
		}

		b, err := json.Marshal(&consumer)
		if err != nil {
			return domain.NewInternalError(
				fmt.Errorf("failed to marshal consumer: %v", err),
				"[domain][consumer][service][ConsumerService][createConsumer]",
			)
		}

		err = srv.fileStorage.Upload(ctx, "consumers/"+consumer.ID.String()+"/consumer.json", b)
		if err != nil {
			return domain.WrapErrorWithStack(err, "[domain][consumer][service][ConsumerService][createConsumer]")
		}

		return nil
	})

	return consumer, err
}
