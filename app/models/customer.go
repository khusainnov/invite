package models

import "time"

type Customer struct {
	Name      *string    `db:"name"`
	Email     *string    `db:"email"`
	CreatedAt *time.Time `db:"created_at"`
}
