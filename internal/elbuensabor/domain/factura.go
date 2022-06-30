package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IFacturaRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, factura Factura) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*Factura, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]Factura, error)
	Update(ctx context.Context, tx *sqlx.Tx, factura Factura) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}
