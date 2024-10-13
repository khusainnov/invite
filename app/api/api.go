package api

import (
	"context"

	"gitlab.com/khusainnov/invite-app/app/models"
)

type EventProcessor interface {
	Process(ctx context.Context, dto *models.EventWithCustomer) error
}

type API struct {
	eventProcessor EventProcessor
}

func New(eventProcessor EventProcessor) *API {
	return &API{eventProcessor: eventProcessor}
}
