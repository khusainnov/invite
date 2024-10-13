package repository

import (
	"context"

	"gitlab.com/khusainnov/invite-app/app/infra/storage"
	"gitlab.com/khusainnov/invite-app/app/models"
)

type CustomerRepo struct{}

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{}
}

func (r *CustomerRepo) Create(
	ctx context.Context,
	q storage.MeasurableQuery,
	customer *models.Customer,
) (int, error) {
	query := `INSERT INTO customer (name, email) VALUES ($1, $2) RETURNING id;`

	var id int
	err := q.QueryRowxContext(ctx, query, customer.Name, customer.Email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CustomerRepo) HasCustomer(
	ctx context.Context,
	q storage.MeasurableQuery,
	gmail string,
) (bool, error) {
	query := `SELECT exists(select 1 FROM customer WHERE email = $1);`

	var exists bool
	err := q.QueryRowxContext(ctx, query, gmail).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, err
}

func (r *CustomerRepo) GetCustomerID(
	ctx context.Context,
	q storage.MeasurableQuery,
	email string,
) (int, error) {
	query := `SELECT id FROM customer WHERE email = $1;`

	var id int
	err := q.QueryRowxContext(ctx, query, email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
