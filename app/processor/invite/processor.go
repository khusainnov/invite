package invite

import (
	"context"

	"gitlab.com/khusainnov/invite-app/app/infra/storage"
	"gitlab.com/khusainnov/invite-app/app/models"
	"go.uber.org/zap"
)

type CustomerRepo interface {
	HasCustomer(ctx context.Context, q storage.MeasurableQuery, gmail string) (bool, error)
	Create(ctx context.Context, q storage.MeasurableQuery, customer *models.Customer) (int, error)
	GetCustomerID(ctx context.Context, q storage.MeasurableQuery, email string) (int, error)
}

type EventRepo interface {
	Create(ctx context.Context, q storage.MeasurableQuery, event *models.Event) error
}

type Processor struct {
	log          *zap.Logger
	query        storage.MeasurableQuery
	customerRepo CustomerRepo
	eventRepo    EventRepo
}

func New(
	log *zap.Logger,
	query storage.MeasurableQuery,
	customerRepo CustomerRepo,
	eventRepo EventRepo,
) *Processor {
	return &Processor{
		log:          log,
		query:        query,
		customerRepo: customerRepo,
		eventRepo:    eventRepo,
	}
}
