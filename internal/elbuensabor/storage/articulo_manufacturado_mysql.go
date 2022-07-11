package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type articuloManufacturadoDB struct {
	ID                   int             `db:"id"`
	TiempoEstimadoCocina sql.NullString  `db:"tiempo_estimado_cocina"`
	Denominacion         sql.NullString  `db:"denominacion"`
	PrecioVenta          sql.NullFloat64 `db:"precio_venta"`
	Imagen               sql.NullString  `db:"imagen"`
}

func (a *articuloManufacturadoDB) toArticuloManufacturado() domain.ArticuloManufacturado {
	return domain.ArticuloManufacturado{
		ID:                   a.ID,
		TiempoEstimadoCocina: database.ToStringP(a.TiempoEstimadoCocina),
		Denominacion:         database.ToStringP(a.Denominacion),
		PrecioVenta:          database.ToFloat64P(a.PrecioVenta),
		Imagen:               database.ToStringP(a.Imagen),
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

func NewMySQLArticuloManufacturadoRepository() *MySQLArticuloManufacturadoRepository {
	return &MySQLArticuloManufacturadoRepository{
		qInsert:     "INSERT INTO articulo_manufacturado (tiempo_estimado_cocina,denominacion, precio_venta, imagen) VALUES (?,?,?,?)",
		qGetByID:    "SELECT id, tiempo_estimado_cocina, denominacion, precio_venta, imagen FROM articulo_manufacturado WHERE id = ?",
		qGetAll:     "SELECT id, tiempo_estimado_cocina, denominacion, precio_venta, imagen FROM articulo_manufacturado",
		qDeleteById: "DELETE FROM articulo_manufacturado WHERE id = ?",
		qUpdate:     "UPDATE articulo_manufacturado SET tiempo_estimado_cocina = COALESCE(?,tiempo_estimado_cocina), denominacion = COALESCE(?,denominacion), precio_venta = COALESCE(?,precio_venta), imagen = COALESCE(?,imagen) WHERE id = ?",
	}
}

func (i *MySQLArticuloManufacturadoRepository) Update(ctx context.Context, tx *sqlx.Tx, artMano domain.ArticuloManufacturado) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, artMano.TiempoEstimadoCocina, artMano.Denominacion, artMano.PrecioVenta, artMano.Imagen, artMano.ID)
	return err
}

func (i *MySQLArticuloManufacturadoRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLArticuloManufacturadoRepository) Insert(ctx context.Context, tx *sqlx.Tx, artMano domain.ArticuloManufacturado) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, artMano.TiempoEstimadoCocina, artMano.Denominacion, artMano.PrecioVenta, artMano.Imagen)
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
