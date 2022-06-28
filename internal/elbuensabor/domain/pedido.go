package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IPedidoRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, pedido Pedido) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*Pedido, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]Pedido, error)
	Update(ctx context.Context, tx *sqlx.Tx, pedido Pedido) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}
