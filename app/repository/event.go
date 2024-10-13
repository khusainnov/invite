package repository

import (
	"context"

	"gitlab.com/khusainnov/invite-app/app/infra/storage"
	"gitlab.com/khusainnov/invite-app/app/models"
)

type EventRepo struct{}

func NewEventRepo() *EventRepo {
	return &EventRepo{}
}

func (r *EventRepo) Create(ctx context.Context, q storage.MeasurableQuery, event *models.Event) error {
	query := `INSERT INTO event (name, member, date) VALUES ($1, $2, $3);`

	_, err := q.ExecContext(ctx, query, event.Name, event.Member, event.Date)
	if err != nil {
		return err
	}

	return nil
}
