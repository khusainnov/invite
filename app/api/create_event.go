package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/khusainnov/invite-app/app/helpers"
	"gitlab.com/khusainnov/invite-app/app/models"
	"gitlab.com/khusainnov/invite-app/specs/event"
)

const (
	deadline = 10
)

func (a *API) Create(r *http.Request, req *event.EventReq, resp *event.EventResp) error {
	ctx, cancelFn := context.WithTimeout(context.TODO(), deadline*time.Second)
	defer cancelFn()

	fmt.Printf("\n%v\n", *req)

	msg, err := buildEventWithCustomer(req)
	if err != nil {
		return fmt.Errorf("failed to parse msg: %w", err)
	}

	if err = a.eventProcessor.Process(ctx, msg); err != nil {
		resp.Message = err.Error()

		return err
	}

	resp.Message = "success"

	return nil
}

func buildEventWithCustomer(req *event.EventReq) (*models.EventWithCustomer, error) {
	eventDate, err := helpers.BuildDatePtr(req.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to parse event_date: %w", err)
	}

	return &models.EventWithCustomer{
		Event: &models.Event{
			Name: &req.Name,
			Date: eventDate,
		},
		Customer: &models.Customer{
			Name:  &req.Customer.Name,
			Email: &req.Customer.Email,
		},
	}, nil
}
