package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type articuloInsumoDB struct {
	ID           int            `db:"id"`
	Denominacion sql.NullString `db:"denominacion"`
	PrecioCompra float64        `db:"precio_compra"`
	PrecioVenta  float64        `db:"precio_venta"`
	StockActual  int            `db:"stock_actual"`
	StockMinimo  int            `db:"stock_minimo"`
	UnidadMedida sql.NullString `db:"unidad_medida"`
	EsInsumo     sql.NullBool   `db:"es_insumo"`
}

func (a *articuloInsumoDB) toArticuloInsumo() domain.ArticuloInsumo {
	return domain.ArticuloInsumo{
		ID:           a.ID,
		Denominacion: database.ToStringP(a.Denominacion),
		PrecioCompra: a.PrecioCompra,
		PrecioVenta:  a.PrecioVenta,
		StockActual:  a.StockActual,
		StockMinimo:  a.StockMinimo,
		UnidadMedida: database.ToStringP(a.UnidadMedida),
		EsInsumo:     database.ToBoolP(a.EsInsumo),
	}
}

type IArticuloInsumoRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloInsumo) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloInsumo, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloInsumo, error)
	Update(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloInsumo) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLArticuloInsumoRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

func NewMySQLArticuloInsumoRepository() *MySQLArticuloInsumoRepository {
	return &MySQLArticuloInsumoRepository{
		qInsert:     "INSERT INTO articulo_insumo (denominacion) VALUES (?)",
		qGetByID:    "SELECT id, denominacion, precio_compra, precio_venta, stock_actual, stock_minimo, unidad_medida, es_insumo FROM articulo_insumo WHERE id = ?",
		qGetAll:     "SELECT id, denominacion, precio_compra, precio_venta, stock_actual, stock_minimo, unidad_medida, es_insumo FROM articulo_insumo",
		qDeleteById: "DELETE FROM articulo_insumo WHERE id = ?",
		qUpdate:     "UPDATE articulo_insumo SET denominacion = COALESCE(?,denominacion), precio_compra = COALESCE(?,precio_compra), precio_venta = COALESCE(?,precio_venta), stock_actual = COALESCE(?,stock_actual), stock_minimo = COALESCE(?,stock_minimo), unidad_medida = COALESCE(?,unidad_medida), es_insumo = COALESCE(?,es_insumo) WHERE id = ?",
	}
}

func (i *MySQLArticuloInsumoRepository) Insert(ctx context.Context, tx *sqlx.Tx, art domain.ArticuloInsumo) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, art.Denominacion, art.PrecioCompra, art.PrecioVenta, art.StockActual, art.StockMinimo, art.UnidadMedida, art.UnidadMedida, art.EsInsumo)
	return err
}

func (i *MySQLArticuloInsumoRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloInsumo, error) {
	query := i.qGetByID
	var articuloInsumo articuloInsumoDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&articuloInsumo)
	if err != nil {
		return nil, err
	}
	artIns := articuloInsumo.toArticuloInsumo()
	return &artIns, nil
}

func (i *MySQLArticuloInsumoRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloInsumo, error) {
	query := i.qGetAll
	articulos := make([]domain.ArticuloInsumo, 0)

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var articulo articuloInsumoDB
		if err := rows.StructScan(&articulo); err != nil {
			return articulos, err
		}
		articulos = append(articulos, articulo.toArticuloInsumo())
	}
	return articulos, nil
}

func (i *MySQLArticuloInsumoRepository) Update(ctx context.Context, tx *sqlx.Tx, art domain.ArticuloInsumo) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, art.Denominacion, art.PrecioCompra, art.PrecioVenta, art.StockActual, art.StockMinimo, art.UnidadMedida, art.UnidadMedida, art.EsInsumo, art.ID)
	return err
}

func (i *MySQLArticuloInsumoRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}
