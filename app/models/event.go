package models

import "time"

type Event struct {
	Name      *string    `db:"name"`
	Member    *int       `db:"member"`
	Date      *time.Time `db:"date"`
	CreatedAt *time.Time `db:"created_at"`
}

type EventWithCustomer struct {
	Event    *Event
	Customer *Customer
}
