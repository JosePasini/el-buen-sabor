package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type IArticuloManufacturadoRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturado) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturado, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturado, error)
	GetAllAvailable(ctx context.Context, tx *sqlx.Tx) ([]*domain.ArticuloManufacturadoAvailable, error)
	Update(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturado) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type articuloManufacturadoDB struct {
	ID                   int             `db:"id"`
	TiempoEstimadoCocina sql.NullInt32   `db:"tiempo_estimado_cocina"`
	Denominacion         sql.NullString  `db:"denominacion"`
	PrecioVenta          sql.NullFloat64 `db:"precio_venta"`
	Imagen               sql.NullString  `db:"imagen"`
}

func (a *articuloManufacturadoDB) toArticuloManufacturado() domain.ArticuloManufacturado {
	return domain.ArticuloManufacturado{
		ID:                   a.ID,
		TiempoEstimadoCocina: database.ToIntP(a.TiempoEstimadoCocina),
		Denominacion:         database.ToStringP(a.Denominacion),
		PrecioVenta:          database.ToFloat64P(a.PrecioVenta),
		Imagen:               database.ToStringP(a.Imagen),
	}
}

type articuloManufacturadoAvailableDB struct {
	CantidadNecesaria     sql.NullInt32   `json:"cantidad_necesaria" db:"cantidad_necesaria"`
	UnidadMedida          sql.NullString  `json:"unidad_medida" db:"unidad_medida"`
	Insumo                sql.NullString  `json:"insumo" db:"insumo"`
	StockActual           sql.NullInt32   `json:"stock_actual" db:"stock_actual"`
	ArticuloManufacturado sql.NullString  `json:"articulo_manufacturado" db:"articulo_manufacturado"`
	TiempoEstimadoCocina  sql.NullInt32   `json:"tiempo_estimado_cocina" db:"tiempo_estimado_cocina"`
	PrecioVenta           sql.NullFloat64 `json:"precio_venta" db:"precio_venta"`
	Disponible            sql.NullBool    `json:"disponible"`
}

func (a *articuloManufacturadoAvailableDB) toArticuloManufacturadoAvailable() domain.ArticuloManufacturadoAvailable {
	return domain.ArticuloManufacturadoAvailable{
		CantidadNecesaria:     database.ToIntP(a.CantidadNecesaria),
		UnidadMedida:          database.ToStringP(a.UnidadMedida),
		Insumo:                database.ToStringP(a.Insumo),
		StockActual:           database.ToIntP(a.StockActual),
		ArticuloManufacturado: database.ToStringP(a.ArticuloManufacturado),
		TiempoEstimadoCocina:  database.ToIntP(a.TiempoEstimadoCocina),
		PrecioVenta:           database.ToFloat64P(a.PrecioVenta),
		Disponible:            database.ToBoolP(a.Disponible),
	}
}

type MySQLArticuloManufacturadoRepository struct {
	qInsert          string
	qGetByID         string
	qGetAll          string
	qGetAllAvailable string
	qDeleteById      string
	qUpdate          string
}

func NewMySQLArticuloManufacturadoRepository() *MySQLArticuloManufacturadoRepository {
	return &MySQLArticuloManufacturadoRepository{
		qInsert:          "INSERT INTO articulo_manufacturado (tiempo_estimado_cocina,denominacion, precio_venta, imagen) VALUES (?,?,?,?)",
		qGetByID:         "SELECT id, tiempo_estimado_cocina, denominacion, precio_venta, imagen FROM articulo_manufacturado WHERE id = ?",
		qGetAll:          "SELECT id, tiempo_estimado_cocina, denominacion, precio_venta, imagen FROM articulo_manufacturado",
		qGetAllAvailable: "select amd.cantidad AS cantidad_necesaria, amd.unidad_medida, ai.denominacion AS insumo, ai.stock_actual, am.denominacion AS articulo_manufacturado, am.tiempo_estimado_cocina, am.precio_venta FROM articulo_manufacturado_detalle amd JOIN articulo_insumo ai ON amd.id_articulo_insumo = ai.id JOIN articulo_manufacturado am ON am.id = amd.id_articulo_manufacturado WHERE ai.es_insumo = true",
		qDeleteById:      "DELETE FROM articulo_manufacturado WHERE id = ?",
		qUpdate:          "UPDATE articulo_manufacturado SET tiempo_estimado_cocina = COALESCE(?,tiempo_estimado_cocina), denominacion = COALESCE(?,denominacion), precio_venta = COALESCE(?,precio_venta), imagen = COALESCE(?,imagen) WHERE id = ?",
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

func (i *MySQLArticuloManufacturadoRepository) GetAllAvailable(ctx context.Context, tx *sqlx.Tx) ([]*domain.ArticuloManufacturadoAvailable, error) {
	query := i.qGetAllAvailable
	articulos := make([]*domain.ArticuloManufacturadoAvailable, 0)
	fmt.Println("Query:", query)
	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var b bool = true
	for rows.Next() {
		var articulo articuloManufacturadoAvailableDB
		if err := rows.StructScan(&articulo); err != nil {
			fmt.Println("err", err)
			return nil, err
		}

		article := articulo.toArticuloManufacturadoAvailable()
		article.Disponible = &b
		articulos = append(articulos, &article)
	}

	fmt.Println("ArticuloManufacturadoAvailable :: ")

	fmt.Printf("%v", articulos)
	return articulos, nil
}
