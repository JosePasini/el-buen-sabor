package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type facturaDB struct {
	ID             int            `db:"id"`
	Fecha          sql.NullTime   `db:"fecha"`
	NumeroFactura  int            `db:"numero_factura"`
	MontoDescuento float64        `db:"monto_descuento"`
	FormaPago      sql.NullString `db:"forma_pago"`
	NumeroTarjeta  sql.NullString `db:"numero_tarjeta"`
	TotalVenta     float64        `db:"total_venta"`
	TotalCosto     float64        `db:"total_costo"`
}

func (i *facturaDB) toFactura() domain.Factura {
	return domain.Factura{
		ID:             i.ID,
		Fecha:          *database.ToTimeP(i.Fecha),
		NumeroFactura:  i.NumeroFactura,
		MontoDescuento: i.MontoDescuento,
		FormaPago:      database.ToStringP(i.FormaPago),
		NumeroTarjeta:  database.ToStringP(i.NumeroTarjeta),
		TotalVenta:     i.TotalVenta,
		TotalCosto:     i.TotalCosto,
	}
}

type MySQLInstrumentoRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

//  fecha (Date)
// número (int) montoDescuento (double) formaPago(String) nroTarjeta(String) totalVenta (double) totalCosto (double)

func NewMySQLInstrumentoRepository() *MySQLInstrumentoRepository {
	return &MySQLInstrumentoRepository{
		qInsert:     "INSERT INTO factura (fecha, numero_factura, monto_descuento, forma_pago, numero_tarjeta, total_venta, total_costo) VALUES (?,?,?,?,?,?,?)",
		qGetByID:    "SELECT id, fecha, numero_factura, monto_descuento, forma_pago, numero_tarjeta, total_venta, total_costo FROM factura WHERE id = ?",
		qGetAll:     "SELECT id, fecha, numero_factura, monto_descuento, forma_pago, numero_tarjeta, total_venta, total_costo FROM factura",
		qDeleteById: "DELETE FROM factura WHERE id = ?",
		qUpdate:     "UPDATE factura SET fecha = COALESCE(?,fecha), numero_factura = COALESCE(?,numero_factura) , monto_descuento = COALESCE(?,monto_descuento), forma_pago = COALESCE(?,forma_pago), numero_tarjeta = COALESCE(?,numero_tarjeta), total_venta = COALESCE(?,total_venta), total_costo = COALESCE(?,total_costo) WHERE id = ?",
	}
}

func (i *MySQLInstrumentoRepository) Update(ctx context.Context, tx *sqlx.Tx, fac domain.Factura) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, fac.Fecha, fac.NumeroFactura, fac.MontoDescuento, fac.FormaPago, fac.NumeroTarjeta, fac.TotalVenta, fac.TotalCosto, fac.ID)
	return err
}

func (i *MySQLInstrumentoRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLInstrumentoRepository) Insert(ctx context.Context, tx *sqlx.Tx, fac domain.Factura) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, fac.Fecha, fac.NumeroFactura, fac.MontoDescuento, fac.FormaPago, fac.NumeroTarjeta, fac.TotalVenta, fac.TotalCosto)
	return err
}

func (i *MySQLInstrumentoRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Factura, error) {
	query := i.qGetByID
	var factura facturaDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&factura)
	if err != nil {
		return nil, err
	}
	fac := factura.toFactura()
	return &fac, nil
}

func (i *MySQLInstrumentoRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Factura, error) {
	query := i.qGetAll
	facturas := make([]domain.Factura, 0)

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var factura facturaDB
		if err := rows.StructScan(&factura); err != nil {
			return facturas, err
		}
		facturas = append(facturas, factura.toFactura())
	}
	return facturas, nil
}
