package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IInstrumentoRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, instrumento Instrumento) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*Instrumento, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]Instrumento, error)
	Update(ctx context.Context, tx *sqlx.Tx, instrumento Instrumento) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}
