package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type IPedidoRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) (int, error)
	InsertDetallePedido(ctx context.Context, tx *sqlx.Tx, detalle_pedido domain.DetallePedido) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Pedido, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Pedido, error)
	Update(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
	UpdateTotal(ctx context.Context, tx *sqlx.Tx, total, id int) error
	DescontarStock(ctx context.Context, tx *sqlx.Tx, idPedido int) (bool, error)
}

type pedidoDB struct {
	ID              int            `db:"id"`
	Estado          int            `db:"estado"`
	HoraEstimadaFin time.Time      `db:"hora_estimada_fin"`
	DetalleEnvio    sql.NullString `db:"detalle_envio"`
	TipoEnvio       int            `db:"tipo_envio"`
	Total           float64        `db:"total"`
	IDDomicicio     int            `db:"id_domicilio"`
	IDCliente       int            `db:"id_cliente"`
}

func (i *pedidoDB) toPedido() domain.Pedido {
	return domain.Pedido{
		ID:              i.ID,
		Estado:          i.Estado,
		HoraEstimadaFin: i.HoraEstimadaFin,
		DetalleEnvio:    database.ToStringP(i.DetalleEnvio),
		TipoEnvio:       i.TipoEnvio,
		Total:           i.Total,
		IDDomicicio:     i.IDDomicicio,
		IDCliente:       i.IDCliente,
	}
}

type MySQLPedidoRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

func NewMySQLPedidoRepository() *MySQLPedidoRepository {
	return &MySQLPedidoRepository{
		qInsert:     "INSERT INTO pedidos (estado, hora_estimada_fin, detalle_envio, tipo_envio, total, id_domicilio, id_cliente) VALUES (?,?,?,?,?,?,?);",
		qGetByID:    "SELECT id, estado, hora_estimada_fin, detalle_envio, tipo_envio, total, id_domicilio, id_cliente FROM pedidos WHERE id = ?",
		qGetAll:     "SELECT id, estado, hora_estimada_fin, detalle_envio, tipo_envio, total, id_domicilio, id_cliente FROM pedidos",
		qDeleteById: "DELETE FROM pedidos WHERE id = ?",
		qUpdate:     "UPDATE pedidos SET estado = COALESCE(?,estado), hora_estimada_fin = COALESCE(?,hora_estimada_fin),detalle_envio = COALESCE(?,detalle_envio),tipo_envio = COALESCE(?,tipo_envio),id_domicilio = COALESCE(?,id_domicilio),id_cliente = COALESCE(?,id_cliente) WHERE id = ?",
	}
}

func (i *MySQLPedidoRepository) Update(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, pedido.Estado, pedido.HoraEstimadaFin, pedido.DetalleEnvio, pedido.TipoEnvio, pedido.Total, pedido.IDDomicicio, pedido.IDCliente, pedido.ID)
	return err
}

func (i *MySQLPedidoRepository) UpdateTotal(ctx context.Context, tx *sqlx.Tx, total, id int) error {
	query := "UPDATE pedidos SET total = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, total, id)
	return err
}

func (i *MySQLPedidoRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLPedidoRepository) Insert(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) (int, error) {
	query := i.qInsert
	sql, err := tx.ExecContext(ctx, query, pedido.Estado, pedido.HoraEstimadaFin, pedido.DetalleEnvio, pedido.TipoEnvio, pedido.Total, pedido.IDDomicicio, pedido.IDCliente)
	if err != nil {
		return 0, err
	}
	idPedido, err := sql.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(idPedido), err
}

func (i *MySQLPedidoRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Pedido, error) {
	query := i.qGetByID
	var pedido pedidoDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&pedido)
	if err != nil {
		return nil, sql.ErrNoRows
	}
	inst := pedido.toPedido()
	return &inst, nil
}

func (i *MySQLPedidoRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Pedido, error) {
	query := i.qGetAll
	pedidos := make([]domain.Pedido, 0)

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pedidoDB pedidoDB
		if err := rows.StructScan(&pedidoDB); err != nil {
			return pedidos, err
		}
		pedidos = append(pedidos, pedidoDB.toPedido())
	}
	return pedidos, nil
}

func (i *MySQLPedidoRepository) InsertDetallePedido(ctx context.Context, tx *sqlx.Tx, det_pedido domain.DetallePedido) error {
	query := "INSERT INTO detalle_pedidos (cantidad, subtotal, id_articulo_manufacturado, id_articulo_insumo, id_pedido) VALUES (?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, query, det_pedido.Cantidad, det_pedido.Subtotal, det_pedido.IdArticuloManufacturado, det_pedido.IdArticuloInsumo, det_pedido.IdPedido)
	return err
}

func (i *MySQLPedidoRepository) DescontarStock(ctx context.Context, tx *sqlx.Tx, idPedido int) (bool, error) {
	var ok bool = true
	type DescontarStockQuery struct {
		IDArticuloInsumo int
		CantidadPedida   int
		CantidadInsumo   int
	}
	var descontarStockList []DescontarStockQuery

	query := `SELECT distinct(ai.id) AS articulo_insumo_id, dp.cantidad AS cantidad_pedida, amd.cantidad AS cantidad_insumo FROM pedidos p
					JOIN detalle_pedidos dp ON dp.id_pedido = p.id
					JOIN articulo_manufacturado am ON am.id = dp.id_articulo_manufacturado
					JOIN articulo_manufacturado_detalle amd ON amd.id_articulo_manufacturado = am.id
					JOIN articulo_insumo ai ON ai.id = amd.id_articulo_insumo
					WHERE p.id = ? AND ai.es_insumo = true`

	rows, err := tx.QueryxContext(ctx, query, idPedido)
	if err != nil {
		return !ok, err
	}
	defer rows.Close()

	for rows.Next() {
		var articulo_insumo_id, cantidad_insumo, cantidad_pedida int
		if err := rows.Scan(&articulo_insumo_id, &cantidad_pedida, &cantidad_insumo); err != nil {
			return !ok, err
		}
		descontarStock := DescontarStockQuery{
			IDArticuloInsumo: articulo_insumo_id,
			CantidadPedida:   cantidad_pedida,
			CantidadInsumo:   cantidad_insumo,
		}
		descontarStockList = append(descontarStockList, descontarStock)
	}

	// Actualizamos el stock del articulo insumo, stock_actual menos la cantidad de insumo utilizado por la cantidad de productos manufacturados pedidos
	//queryDescontarStock := "UPDATE articulo_insumo SET stock_actual = (stock_actual - CantidadInsumo * CantidadPedida) WHERE IDArticuloInsumo = ?"

	fmt.Println("?", descontarStockList)
	return false, err

	for _, des := range descontarStockList {
		query_slices := []string{"UPDATE articulo_insumo SET stock_actual = (stock_actual - (", strconv.Itoa(des.CantidadInsumo),
			" * ", strconv.Itoa(des.CantidadPedida), ")) WHERE id = ", strconv.Itoa(des.IDArticuloInsumo)}
		queryDescontarStockOk := strings.Join(query_slices, "")
		fmt.Println("query ::", queryDescontarStockOk)
		_, err := tx.ExecContext(ctx, queryDescontarStockOk)
		if err != nil {
			return !ok, err
		}
	}

	// Actualizo el 'estado' del pedido a 2 :: Estado 'aceptado'
	queryActualizarEstadoPedido := "UPDATE pedidos SET estado = 2 WHERE id = ?"
	_, err = tx.ExecContext(ctx, queryActualizarEstadoPedido, idPedido)
	if err != nil {
		return !ok, err
	}

	fmt.Println("Ok", ok)
	return ok, err
}
