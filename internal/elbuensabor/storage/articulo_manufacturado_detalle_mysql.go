package storage

import (
	"context"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type articuloManufacturadoDetalleDB struct {
	ID                      int `db:"id"`
	Cantidad                int `db:"cantidad"`
	IDArticuloManufacturado int `db:"id_articulo_manufacturado"`
	IDArticuloInsumo        int `db:"id_articulo_insumo"`
}

func (a *articuloManufacturadoDetalleDB) toArticuloManufacturadoDetalle() domain.ArticuloManufacturadoDetalle {
	return domain.ArticuloManufacturadoDetalle{
		ID:                      a.ID,
		Cantidad:                a.Cantidad,
		IDArticuloManufacturado: a.IDArticuloManufacturado,
		IDArticuloInsumo:        a.IDArticuloInsumo,
	}
}

type IArticuloManufacturadoDetalleRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturadoDetalle, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturadoDetalle, error)
	Update(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLArticuloManufacturadoDetalleRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

func NewMySQLArticuloManufacturadoDetalleRepository() *MySQLArticuloManufacturadoDetalleRepository {
	return &MySQLArticuloManufacturadoDetalleRepository{
		qInsert:     "INSERT INTO articulo_manufacturado_detalle (cantidad, id_articulo_insumo, id_articulo_manufacturado) VALUES (?,?,?)",
		qGetByID:    "SELECT id, cantidad, id_articulo_manufacturado, id_articulo_insumo FROM articulo_manufacturado_detalle WHERE id = ?",
		qGetAll:     "SELECT id, cantidad, id_articulo_manufacturado, id_articulo_insumo FROM articulo_manufacturado_detalle",
		qDeleteById: "DELETE FROM articulo_manufacturado_detalle WHERE id = ?",
		qUpdate:     "UPDATE articulo_manufacturado_detalle SET cantidad = COALESCE(?,cantidad), id_articulo_manufacturado = COALESCE(?,id_articulo_manufacturado),id_articulo_insumo = COALESCE(?,id_articulo_insumo) WHERE id = ?",
	}
}
func (i *MySQLArticuloManufacturadoDetalleRepository) Update(ctx context.Context, tx *sqlx.Tx, artManu domain.ArticuloManufacturadoDetalle) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, artManu.Cantidad, artManu.IDArticuloManufacturado, artManu.IDArticuloInsumo, artManu.ID)
	return err
}

func (i *MySQLArticuloManufacturadoDetalleRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLArticuloManufacturadoDetalleRepository) Insert(ctx context.Context, tx *sqlx.Tx, artManu domain.ArticuloManufacturadoDetalle) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, artManu.Cantidad, artManu.IDArticuloManufacturado, artManu.IDArticuloInsumo)
	return err
}

func (i *MySQLArticuloManufacturadoDetalleRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturadoDetalle, error) {
	query := i.qGetByID
	var articuloManufacturadoDetalle articuloManufacturadoDetalleDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&articuloManufacturadoDetalle)
	if err != nil {
		return nil, err
	}
	artManufacturadoDetalle := articuloManufacturadoDetalle.toArticuloManufacturadoDetalle()
	return &artManufacturadoDetalle, nil
}

func (i *MySQLArticuloManufacturadoDetalleRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturadoDetalle, error) {
	query := i.qGetAll
	articulosManufacturadoDetalles := make([]domain.ArticuloManufacturadoDetalle, 0)

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var articuloManufacturadoDetalle articuloManufacturadoDetalleDB
		if err := rows.StructScan(&articuloManufacturadoDetalle); err != nil {
			return articulosManufacturadoDetalles, err
		}
		articulosManufacturadoDetalles = append(articulosManufacturadoDetalles, articuloManufacturadoDetalle.toArticuloManufacturadoDetalle())
	}
	return articulosManufacturadoDetalles, nil
}
