package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IFacturaRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, factura Factura) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*Factura, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]Factura, error)
	RecaudacionesDiarias(ctx context.Context, tx *sqlx.Tx, fecha string) ([]Recaudaciones, error)
	RecaudacionesMensuales(ctx context.Context, tx *sqlx.Tx, month, year string) ([]Recaudaciones, error)
	RecaudacionesPeriodoTiempo(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]RecaudacionesResponse, error)
	ObtenerGanancias(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]Ganancias, error)
	Update(ctx context.Context, tx *sqlx.Tx, factura Factura) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}
