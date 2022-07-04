package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type articuloManufacturadoDetalleDB struct {
	ID           int            `db:"id"`
	UnidadMedida sql.NullString `db:"unidad_medida"`
}

func (a *articuloManufacturadoDetalleDB) toArticuloManufacturadoDetalle() domain.ArticuloManufacturadoDetalle {
	return domain.ArticuloManufacturadoDetalle{
		ID:           a.ID,
		UnidadMedida: *database.ToStringP(a.UnidadMedida),
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
		qInsert:     "INSERT INTO articulo_manufacturado_detalle (unidad_medida) VALUES (?)",
		qGetByID:    "SELECT id, unidad_medida FROM articulo_manufacturado_detalle WHERE id = ?",
		qGetAll:     "SELECT id, unidad_medida FROM articulo_manufacturado_detalle",
		qDeleteById: "DELETE FROM articulo_manufacturado_detalle WHERE id = ?",
		qUpdate:     "UPDATE articulo_manufacturado_detalle SET unidad_medida = COALESCE(?,unidad_medida) WHERE id = ?",
	}
}
func (i *MySQLArticuloManufacturadoDetalleRepository) Update(ctx context.Context, tx *sqlx.Tx, artManufacturado domain.ArticuloManufacturadoDetalle) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, artManufacturado.UnidadMedida, artManufacturado.ID)
	return err
}

func (i *MySQLArticuloManufacturadoDetalleRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLArticuloManufacturadoDetalleRepository) Insert(ctx context.Context, tx *sqlx.Tx, artManufacturado domain.ArticuloManufacturadoDetalle) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, artManufacturado.UnidadMedida)
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
