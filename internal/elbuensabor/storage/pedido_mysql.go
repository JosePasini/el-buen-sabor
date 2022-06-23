package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

/*
   `id_cliente` INT,
   `fecha` DATETIME,
   `domicilio_envio` VARCHAR(255),
   `detalle_envio` VARCHAR(255),
   `delivery` BOOLEAN,
   `metodo_pago` ENUM('efectivo','mercadopago'),
*/

type pedidoDB struct {
	ID             int            `db:"id"`
	IDCliente      int            `db:"id_cliente"`
	Fecha          time.Time      `db:"fecha"`
	DomicilioEnvio sql.NullString `db:"domicilio_envio"`
	DetalleEnvio   sql.NullString `db:"detalle_envio"`
	Delivery       sql.NullBool   `db:"delivery"`
	MetodoPago     sql.NullString `db:"metodo_pago"`
}

func (i *pedidoDB) toPedido() domain.Pedido {
	return domain.Pedido{
		ID:             i.ID,
		IDCliente:      i.IDCliente,
		Fecha:          i.Fecha,
		DomicilioEnvio: database.ToStringP(i.DomicilioEnvio),
		DetalleEnvio:   database.ToStringP(i.DetalleEnvio),
		Delivery:       database.ToBoolP(i.Delivery),
		MetodoPago:     database.ToStringP(i.MetodoPago),
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
		qInsert:     "INSERT INTO pedidos (id_cliente, domicilio_envio, detalle_envio, delivery, metodo_pago, fecha) VALUES (?,?,?,?,?, now());",
		qGetByID:    "SELECT id,id_cliente, domicilio_envio, detalle_envio, delivery, metodo_pago FROM pedidos WHERE id = ?",
		qGetAll:     "SELECT id,id_cliente, domicilio_envio, detalle_envio, delivery, metodo_pago FROM pedidos",
		qDeleteById: "DELETE FROM pedidos WHERE id = ?",
		qUpdate:     "UPDATE pedidos SET id_cliente = COALESCE(?,id_cliente), domicilio_envio = COALESCE(?,domicilio_envio),detalle_envio = COALESCE(?,detalle_envio),delivery = COALESCE(?,delivery),metodo_pago = COALESCE(?,metodo_pago) WHERE id = ?",
	}
}

func (i *MySQLPedidoRepository) Update(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, pedido.IDCliente, pedido.DomicilioEnvio, pedido.DetalleEnvio, pedido.Delivery, pedido.MetodoPago, pedido.ID)
	return err
}

func (i *MySQLPedidoRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLPedidoRepository) Insert(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, pedido.IDCliente, pedido.DomicilioEnvio, pedido.DetalleEnvio, pedido.Delivery, pedido.MetodoPago)
	return err
}

func (i *MySQLPedidoRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Pedido, error) {
	query := i.qGetByID
	var pedido pedidoDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&pedido)
	if err != nil {
		return nil, err
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
