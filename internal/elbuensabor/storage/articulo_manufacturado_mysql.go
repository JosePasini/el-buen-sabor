package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type articuloManufacturadoDB struct {
	ID           int            `db:"id"`
	Denominacion sql.NullString `db:"denominacion"`
}

func (a *articuloManufacturadoDB) toArticuloManufacturado() domain.ArticuloManufacturado {
	return domain.ArticuloManufacturado{
		ID:           a.ID,
		Denominacion: database.ToStringP(a.Denominacion),
	}
}

type IArticuloManufacturadoRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturado) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturado, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturado, error)
	Update(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturado) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLArticuloManufacturadoRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

//  fecha (Date)
// número (int) montoDescuento (double) formaPago(String) nroTarjeta(String) totalVenta (double) totalCosto (double)

func NewMySQLArticuloManufacturadoRepository() *MySQLArticuloManufacturadoRepository {
	return &MySQLArticuloManufacturadoRepository{
		qInsert:     "INSERT INTO articulo_manufacturado (denominacion) VALUES (?)",
		qGetByID:    "SELECT id, denominacion FROM articulo_manufacturado WHERE id = ?",
		qGetAll:     "SELECT id, denominacion FROM articulo_manufacturado",
		qDeleteById: "DELETE FROM articulo_manufacturado WHERE id = ?",
		qUpdate:     "UPDATE articulo_manufacturado SET denominacion = COALESCE(?,denominacion) WHERE id = ?",
	}
}

func (i *MySQLArticuloManufacturadoRepository) Update(ctx context.Context, tx *sqlx.Tx, artManufacturado domain.ArticuloManufacturado) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, artManufacturado.Denominacion, artManufacturado.ID)
	return err
}

func (i *MySQLArticuloManufacturadoRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLArticuloManufacturadoRepository) Insert(ctx context.Context, tx *sqlx.Tx, artManufacturado domain.ArticuloManufacturado) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, artManufacturado.Denominacion)
	return err
}

func (i *MySQLArticuloManufacturadoRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturado, error) {
	query := i.qGetByID
	var articulo articuloManufacturadoDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&articulo)
	if err != nil {
		return nil, err
	}
	artManufacturado := articulo.toArticuloManufacturado()
	return &artManufacturado, nil
}

func (i *MySQLArticuloManufacturadoRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturado, error) {
	query := i.qGetAll
	articulos := make([]domain.ArticuloManufacturado, 0)

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var articulo articuloManufacturadoDB
		if err := rows.StructScan(&articulo); err != nil {
			return articulos, err
		}
		articulos = append(articulos, articulo.toArticuloManufacturado())
	}
	return articulos, nil
}
