package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//DB :: Interface DB
type DB interface {
	WithTransaction(ctx context.Context, txFunc func(tx *sqlx.Tx) error) (err error)
	Close() error
}

//NoopDB :: Structure noop db
type NoopDB struct{}

//WithTransaction :: Manejo transaction DB
func (db NoopDB) WithTransaction(_ context.Context, f func(*sqlx.Tx) error) (err error) {
	tx := &sqlx.Tx{}
	return f(tx)
}

//Close :: Close DB
func (db NoopDB) Close() error {
	return nil
}
