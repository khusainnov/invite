package storage

import (
	"github.com/jmoiron/sqlx"
)

type MeasurableQuery interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
}
