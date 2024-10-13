package invite

import (
	"context"
	"fmt"

	"gitlab.com/khusainnov/invite-app/app/models"
	"go.uber.org/zap"
)

func (p *Processor) Process(ctx context.Context, dto *models.EventWithCustomer) error {
	log := p.log.With(zap.String("handler", "inviteapp.Processor.Create"))

	log.Info("create", zap.String("event_name", *dto.Event.Name))

	id, err := p.processCustomer(ctx, log, dto.Customer)
	if err != nil {
		return err
	}

	dto.Event.Member = &id

	err = p.processEvent(ctx, log, dto.Event)
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) processCustomer(ctx context.Context, log *zap.Logger, customer *models.Customer) (int, error) {
	hasCustomer, err := p.customerRepo.HasCustomer(ctx, p.query, *customer.Email)
	if err != nil {
		return 0, err
	}

	var id int
	if hasCustomer {
		id, err = p.customerRepo.GetCustomerID(ctx, p.query, *customer.Email)
		if err != nil {
			log.Error("failed to get customer_id", zap.Error(err))

			return 0, err
		}

		return id, nil
	}

	log.Info("customer doesn't exists, creating new")

	id, err = p.customerRepo.Create(ctx, p.query, customer)
	if err != nil {
		log.Error("failed to create customer", zap.Error(err))

		return 0, fmt.Errorf("failed to create customer: %w", err)
	}

	return id, nil
}

func (p *Processor) processEvent(ctx context.Context, log *zap.Logger, event *models.Event) error {
	log.Info("creating event")

	if err := p.eventRepo.Create(ctx, p.query, event); err != nil {
		log.Error("failed to create event", zap.Int("member_id", *event.Member))

		return fmt.Errorf("failed to create event: %w", err)
	}

	return nil
}
