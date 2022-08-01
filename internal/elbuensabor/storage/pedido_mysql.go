package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

var (
	StockInsuficiente = errors.New("stock insuficiente")
)

type IPedidoRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido, minutosDemoraCocina int) (int, error)
	//InsertDetallePedido(ctx context.Context, tx *sqlx.Tx, detalle_pedido domain.DetallePedido) error
	InsertDetallePedido(ctx context.Context, tx *sqlx.Tx, carrito_completo domain.CarritoCompleto) error
	GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Pedido, error)
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Pedido, error)
	GetAllDetallePedidoByIDPedido(ctx context.Context, tx *sqlx.Tx, idPedido int) ([]domain.DetallePedidoResponse, error)
	GetAllPedidosByIDCliente(ctx context.Context, tx *sqlx.Tx, idCliente int) ([]domain.Pedido, error)
	GetRankingDePedidosPorCliente(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]domain.PedidosPorCliente, error)
	GetCostoTotalByPedido(ctx context.Context, tx *sqlx.Tx, idPedido int) (float64, error)
	Update(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
	CancelarPedido(ctx context.Context, tx *sqlx.Tx, idPedido int) error
	UpdateTotal(ctx context.Context, tx *sqlx.Tx, total, id int) error
	DescontarStock(ctx context.Context, tx *sqlx.Tx, idPedido, estado int) (bool, error)
	UpdateEstadoPedido(ctx context.Context, tx *sqlx.Tx, estado, IDPedido int) error
	RankingComidasMasPedidas(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]domain.RankingComidasMasPedidas, error)
	VerificarStockBebidas(ctx context.Context, tx *sqlx.Tx, idArticulo, amount int) (bool, error)
	VerificarStockManufacturado(ctx context.Context, tx *sqlx.Tx, idArticulo, amount int) (bool, error)
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

type detallePedidoDB struct {
	IDPedido     int            `db:"id_pedido"`
	Cantidad     int            `db:"cantidad"`
	Subtotal     float64        `db:"subtotal"`
	Denominacion sql.NullString `db:"denominacion"`
	Imagen       sql.NullString `db:"imagen"`
}

func (i *detallePedidoDB) toDetallePedido() domain.DetallePedidoResponse {
	return domain.DetallePedidoResponse{
		IDPedido:     i.IDPedido,
		Cantidad:     i.Cantidad,
		Subtotal:     i.Subtotal,
		Denominacion: database.ToStringP(i.Denominacion),
		Imagen:       database.ToStringP(i.Imagen),
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
		//qInsert:     "INSERT INTO pedidos (estado, hora_estimada_fin, detalle_envio, tipo_envio, total, id_domicilio, id_cliente) VALUES (?,DATE_ADD(?,INTERVAL ? MINUTE),?,?,?,?,?);",
		// El insert de arriba es el correcto, pero como la zona horaria de heroku es en USA, tengo que restarle 3 horas. select DATE_ADD(DATE_SUB(NOW(), INTERVAL 3 HOUR), INTERVAL 30 MINUTE);
		qInsert:     "INSERT INTO pedidos (estado, hora_estimada_fin, detalle_envio, tipo_envio, total, id_domicilio, id_cliente) VALUES (?,DATE_ADD(DATE_SUB(NOW(), INTERVAL 3 HOUR), INTERVAL ? MINUTE),?,?,?,?,?);",
		qGetByID:    "SELECT id, estado, hora_estimada_fin, detalle_envio, tipo_envio, total, id_domicilio, id_cliente FROM pedidos WHERE id = ?",
		qGetAll:     "SELECT id, estado, hora_estimada_fin, detalle_envio, tipo_envio, total, id_domicilio, id_cliente FROM pedidos",
		qDeleteById: "DELETE FROM pedidos WHERE id = ?",
		qUpdate:     "UPDATE pedidos SET estado = COALESCE(?,estado), hora_estimada_fin = COALESCE(?,hora_estimada_fin),detalle_envio = COALESCE(?,detalle_envio),tipo_envio = COALESCE(?,tipo_envio),id_domicilio = COALESCE(?,id_domicilio),id_cliente = COALESCE(?,id_cliente) WHERE id = ?",
	}
}

func (i *MySQLPedidoRepository) Update(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, pedido.Estado, pedido.HoraEstimadaFin, pedido.DetalleEnvio, pedido.TipoEnvio, pedido.IDDomicicio, pedido.IDCliente, pedido.ID)
	return err
}

func (i *MySQLPedidoRepository) UpdateEstadoPedido(ctx context.Context, tx *sqlx.Tx, estado, IDPedido int) error {
	query := "UPDATE pedidos SET estado = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, estado, IDPedido)
	fmt.Println("estado:", estado)
	fmt.Println("IDPedido:", IDPedido)
	fmt.Println("err:", err)
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

func (i *MySQLPedidoRepository) Insert(ctx context.Context, tx *sqlx.Tx, pedido domain.Pedido, minutosDemoraCocina int) (int, error) {
	query := i.qInsert
	sql, err := tx.ExecContext(ctx, query, pedido.Estado, minutosDemoraCocina, pedido.DetalleEnvio, pedido.TipoEnvio, pedido.Total, pedido.IDDomicicio, pedido.IDCliente)
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

func (i *MySQLPedidoRepository) GetAllPedidosByIDCliente(ctx context.Context, tx *sqlx.Tx, idCliente int) ([]domain.Pedido, error) {
	queryGetPedidosByCliente := `SELECT * FROM pedidos
    	WHERE id_cliente = ?;`

	pedidos := make([]domain.Pedido, 0)

	rows, err := tx.QueryxContext(ctx, queryGetPedidosByCliente, idCliente)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pedido pedidoDB
		if err := rows.StructScan(&pedido); err != nil {
			return pedidos, err
		}
		pedidos = append(pedidos, pedido.toPedido())
	}

	return pedidos, nil
}

func (i *MySQLPedidoRepository) GetAllDetallePedidoByIDPedido(ctx context.Context, tx *sqlx.Tx, idPedido int) ([]domain.DetallePedidoResponse, error) {

	queryGetManufacturados := `SELECT dp.id_pedido, dp.cantidad, dp.subtotal, am.denominacion, am.imagen FROM detalle_pedidos dp
		JOIN articulo_manufacturado am ON am.id = dp.id_articulo_manufacturado
    	WHERE dp.id_pedido = ?;`

	queryGetBebidas := `SELECT dp.id_pedido,dp.cantidad, dp.subtotal,ai.denominacion, ai.imagen FROM detalle_pedidos dp
		JOIN articulo_insumo ai ON ai.id = dp.id_articulo_insumo
		WHERE dp.id_pedido = ? AND ai.es_insumo = false;`

	pedidos := make([]domain.DetallePedidoResponse, 0)

	rows, err := tx.QueryxContext(ctx, queryGetManufacturados, idPedido)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pedido detallePedidoDB
		if err := rows.StructScan(&pedido); err != nil {
			return pedidos, err
		}
		pedidos = append(pedidos, pedido.toDetallePedido())
	}

	rows, err = tx.QueryxContext(ctx, queryGetBebidas, idPedido)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println("antes:::")

	for rows.Next() {
		var pedido detallePedidoDB
		if err := rows.StructScan(&pedido); err != nil {
			return pedidos, err
		}
		pedidos = append(pedidos, pedido.toDetallePedido())
	}
	return pedidos, nil
}

func (i *MySQLPedidoRepository) GetRankingDePedidosPorCliente(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]domain.PedidosPorCliente, error) {
	queryGetManufacturados := `SELECT count(p.id) AS cantidad_pedidos, p.id_cliente, SUM(total) AS total FROM pedidos p
	WHERE p.estado = 5
    AND p.hora_estimada_fin BETWEEN ? AND ?
    GROUP BY p.id_cliente
	ORDER BY cantidad_pedidos DESC;`

	pedidos := make([]domain.PedidosPorCliente, 0)

	rows, err := tx.QueryxContext(ctx, queryGetManufacturados, desde, hasta)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cantidad_pedidos, id_cliente int
		var total float64
		if err := rows.Scan(&cantidad_pedidos, &id_cliente, &total); err != nil {
			return pedidos, err
		}
		pedByClient := domain.PedidosPorCliente{
			CantidadPedidos: cantidad_pedidos,
			IDCliente:       id_cliente,
			Total:           total,
		}
		pedidos = append(pedidos, pedByClient)
	}
	return pedidos, nil
}

// func (i *MySQLPedidoRepository) InsertDetallePedido(ctx context.Context, tx *sqlx.Tx, det_pedido domain.DetallePedido) error {
// 	query := "INSERT INTO detalle_pedidos (cantidad, subtotal, id_articulo_manufacturado, id_articulo_insumo, id_pedido) VALUES (?,?,?,?,?)"
// 	_, err := tx.ExecContext(ctx, query, det_pedido.Cantidad, det_pedido.Subtotal, det_pedido.IdArticuloManufacturado, det_pedido.IdArticuloInsumo, det_pedido.IdPedido)
// 	return err
// }

func (i *MySQLPedidoRepository) InsertDetallePedido(ctx context.Context, tx *sqlx.Tx, carrito domain.CarritoCompleto) error {
	//query := "INSERT INTO detalle_pedidos (cantidad, subtotal, id_articulo_manufacturado, id_articulo_insumo, id_pedido) VALUES (?,?,?,?,?)"
	if carrito.EsBebida {
		query := "INSERT INTO detalle_pedidos (cantidad, subtotal, id_articulo_insumo, id_pedido) VALUES (?,?,?,?)"
		_, err := tx.ExecContext(ctx, query, carrito.Cantidad, carrito.SubTotal, carrito.ID, carrito.IDPedido)
		return err
	} else {
		query := "INSERT INTO detalle_pedidos (cantidad, subtotal, id_articulo_manufacturado, id_pedido) VALUES (?,?,?,?)"
		_, err := tx.ExecContext(ctx, query, carrito.Cantidad, carrito.SubTotal, carrito.ID, carrito.IDPedido)
		return err
	}
}

func (i *MySQLPedidoRepository) RankingComidasMasPedidas(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]domain.RankingComidasMasPedidas, error) {
	query := `SELECT p.id AS id_pedido, SUM(dp.cantidad) AS veces_pedida, dp.id_articulo_manufacturado, am.denominacion FROM pedidos p 
				JOIN detalle_pedidos dp on dp.id_pedido = p.id
    			JOIN articulo_manufacturado am on am.id = dp.id_articulo_manufacturado
    			WHERE id_articulo_manufacturado IS NOT NULL
    			AND p.hora_estimada_fin BETWEEN ? AND ?
				AND p.estado = 5
    			GROUP BY id_articulo_manufacturado
    			ORDER BY veces_pedida desc;`

	rankingComidas := make([]domain.RankingComidasMasPedidas, 0)

	rows, err := tx.QueryxContext(ctx, query, desde, hasta)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id_pedido, veces_pedida, id_articulo_manufacturado int
		var denominacion string
		if err := rows.Scan(&id_pedido, &veces_pedida, &id_articulo_manufacturado, &denominacion); err != nil {
			return nil, err
		}
		rankComidas := domain.RankingComidasMasPedidas{
			IDPedido:                id_pedido,
			VecesPedida:             veces_pedida,
			IDArticuloManufacturado: id_articulo_manufacturado,
			Denominacion:            denominacion,
		}
		rankingComidas = append(rankingComidas, rankComidas)
	}
	return rankingComidas, nil
}

func (i *MySQLPedidoRepository) DescontarStock(ctx context.Context, tx *sqlx.Tx, idPedido, estado int) (bool, error) {
	var err error
	var ok bool

	// Descontar stock de los insumos utilizados en los artículos manufacturados
	ok, err = DescontarStockManufacturado(ctx, tx, idPedido)
	if err != nil {
		return !ok, err
	}

	// Descontar stock bebidas
	ok, err = DescontarStockBebidas(ctx, tx, idPedido)
	if err != nil {
		return !ok, err
	}

	// Actualizo el 'estado' del pedido a 2 :: Estado 'aceptado'
	var queryActualizarEstadoPedido string
	if estado == domain.PENDIENTE_APROBACION {
		queryActualizarEstadoPedido = "UPDATE pedidos SET estado = 2 WHERE id = ?"
		_, err = tx.ExecContext(ctx, queryActualizarEstadoPedido, idPedido)
		if err != nil {
			return !ok, err
		}
	}

	fmt.Println("Ok", ok)
	return ok, err
}

func (i *MySQLPedidoRepository) GetCostoTotalByPedido(ctx context.Context, tx *sqlx.Tx, idPedido int) (float64, error) {
	var err error
	var costo_total_final float64
	// query para obtener el COSTO TOTAL de bebidas por pedido
	queryCostoBebidas := `SELECT SUM(precio_compra * dp.cantidad) AS costo_total FROM detalle_pedidos dp
			JOIN articulo_insumo ai ON ai.id = dp.id_articulo_insumo
   			WHERE dp.id_pedido = ?;`

	// query para obtener el COSTO TOTAL de manufacturados por pedido
	queryCostoManufacturados := `SELECT SUM(precio_compra * dp.cantidad) AS costo_total FROM detalle_pedidos dp
			JOIN articulo_manufacturado am ON am.id = dp.id_articulo_manufacturado
    		JOIN articulo_manufacturado_detalle amd ON amd.id_articulo_manufacturado = am.id
			JOIN articulo_insumo ai ON ai.id = amd.id_articulo_insumo
    		WHERE dp.id_pedido = ?;`

	rows, err := tx.QueryxContext(ctx, queryCostoBebidas, idPedido)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		//var costo_total float64
		var costo_total sql.NullFloat64
		if err := rows.Scan(&costo_total); err != nil {
			return 0, err
		}
		costo_total_final = costo_total.Float64
	}

	rows, err = tx.QueryxContext(ctx, queryCostoManufacturados, idPedido)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		//var costo_total float64
		var costo_total sql.NullFloat64
		if err := rows.Scan(&costo_total); err != nil {
			return 0, err
		}
		costo_total_final = costo_total.Float64 + costo_total_final
	}
	fmt.Println("costo_total_final", costo_total_final)
	return costo_total_final, err
}

func DescontarStockBebidas(ctx context.Context, tx *sqlx.Tx, idPedido int) (bool, error) {
	var ok bool = true
	type DescontarStockBebidas struct {
		StockActual int
		ArticuloID  int
	}
	var descontarStockList []DescontarStockBebidas

	// Descontar Stock articulos NO insumo :: Cervezas, gaseosas, etc.
	queryArticuloInsumo := `SELECT (ai.stock_actual - dp.cantidad) AS stock_actual, ai.id FROM pedidos p
					JOIN detalle_pedidos dp ON dp.id_pedido = p.id
					JOIN articulo_insumo ai ON ai.id = dp.id_articulo_insumo
					WHERE p.id = ? AND ai.es_insumo = false;`

	rows, err := tx.QueryxContext(ctx, queryArticuloInsumo, idPedido)
	if err != nil {
		return !ok, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock_actual, articulo_id int
		if err := rows.Scan(&stock_actual, &articulo_id); err != nil {
			return !ok, err
		}
		descontarStock := DescontarStockBebidas{
			StockActual: stock_actual,
			ArticuloID:  articulo_id,
		}
		descontarStockList = append(descontarStockList, descontarStock)
	}
	if err != nil {
		return !ok, err
	}
	fmt.Println("descontarStockList ::", descontarStockList)

	for _, des := range descontarStockList {
		stockActual := strconv.Itoa(des.StockActual)
		articuloID := strconv.Itoa(des.ArticuloID)
		fmt.Println("stockActual", stockActual)
		if des.StockActual < 0 {
			return !ok, StockInsuficiente
		}
		query_slices := []string{"UPDATE articulo_insumo SET stock_actual = ", stockActual, " WHERE id = ", articuloID}
		queryDescontarStock := strings.Join(query_slices, "")
		fmt.Println("query 2 ::", queryDescontarStock)
		_, err := tx.ExecContext(ctx, queryDescontarStock)
		if err != nil {
			return ok, err
		}
	}
	if err != nil {
		return !ok, err
	}
	return ok, nil
}

func (i *MySQLPedidoRepository) VerificarStockBebidas(ctx context.Context, tx *sqlx.Tx, idArticulo, amount int) (bool, error) {
	var ok bool = true
	queryArticuloInsumo := `SELECT IF( (( articulo_insumo.stock_actual - ? ) >= 0), true, false) from articulo_insumo WHERE id = ?;`

	rows, err := tx.QueryxContext(ctx, queryArticuloInsumo, amount, idArticulo)
	if err != nil {
		return !ok, err
	}
	defer rows.Close()

	for rows.Next() {
		var ok bool
		if err := rows.Scan(&ok); err != nil {
			return ok, err
		}
		fmt.Println("Ok:", ok)
		if !ok {
			return ok, StockInsuficiente
		}
	}
	return ok, nil
}

func DescontarStockManufacturado(ctx context.Context, tx *sqlx.Tx, idPedido int) (bool, error) {
	var ok bool = true
	type DescontarStockInsumosQuery struct {
		IDArticuloInsumo int
		CantidadPedida   int
		CantidadInsumo   int
	}

	var descontarStockList []DescontarStockInsumosQuery

	queryArticuloManufacturado := `SELECT ai.id AS articulo_insumo_id, dp.cantidad AS cantidad_pedida, amd.cantidad AS cantidad_insumo FROM pedidos p
					JOIN detalle_pedidos dp ON dp.id_pedido = p.id
					JOIN articulo_manufacturado am ON am.id = dp.id_articulo_manufacturado
					JOIN articulo_manufacturado_detalle amd ON amd.id_articulo_manufacturado = am.id
					JOIN articulo_insumo ai ON ai.id = amd.id_articulo_insumo
					WHERE p.id = ? AND ai.es_insumo = true`

	rows, err := tx.QueryxContext(ctx, queryArticuloManufacturado, idPedido)
	if err != nil {
		return !ok, err
	}
	defer rows.Close()

	for rows.Next() {
		var articulo_insumo_id, cantidad_insumo, cantidad_pedida int
		if err := rows.Scan(&articulo_insumo_id, &cantidad_pedida, &cantidad_insumo); err != nil {
			return !ok, err
		}
		descontarStock := DescontarStockInsumosQuery{
			IDArticuloInsumo: articulo_insumo_id,
			CantidadPedida:   cantidad_pedida,
			CantidadInsumo:   cantidad_insumo,
		}
		descontarStockList = append(descontarStockList, descontarStock)
	}

	// Actualizamos el stock del articulo insumo, stock_actual menos la cantidad de insumo utilizado por la cantidad de productos manufacturados pedidos
	//queryDescontarStock := "UPDATE articulo_insumo SET stock_actual = (stock_actual - CantidadInsumo * CantidadPedida) WHERE IDArticuloInsumo = ?"

	for _, des := range descontarStockList {
		// query_slices := []string{"UPDATE articulo_insumo SET stock_actual = (stock_actual - (", strconv.Itoa(des.CantidadInsumo),
		// 	" * ", strconv.Itoa(des.CantidadPedida), ")) WHERE id = ", strconv.Itoa(des.IDArticuloInsumo)}
		cantInsumo := strconv.Itoa(des.CantidadInsumo)
		cantPedida := strconv.Itoa(des.CantidadPedida)
		artInsumo := strconv.Itoa(des.IDArticuloInsumo)

		queryOk := []string{"SELECT IF( ((stock_actual - (", cantInsumo,
			" * ", cantPedida, ")) > 0), true, false ) FROM articulo_insumo WHERE id = ", artInsumo}
		queryOkAux := strings.Join(queryOk, "")
		fmt.Println("query:", queryOkAux)
		//queryOk := select IF( ((stock_actual - (100 * 55) ) > 0), (stock_actual - 100 * 55), true ) from articulo_insumo WHERE id = 94;
		rows, err := tx.QueryxContext(ctx, queryOkAux)
		if err != nil {
			return !ok, err
		}
		defer rows.Close()

		for rows.Next() {
			var ok bool
			if err := rows.Scan(&ok); err != nil {
				return ok, err
			}
			fmt.Println("Ok:", ok)
			if !ok {
				return ok, StockInsuficiente
			}
		}

		query_slices := []string{"UPDATE articulo_insumo SET stock_actual = IF( ((stock_actual - (", cantInsumo,
			" * ", cantPedida, ")) > 0), (stock_actual - ", cantInsumo, " * ", cantPedida, ") , stock_actual ) WHERE id = ", artInsumo}
		queryDescontarStockOk := strings.Join(query_slices, "")
		fmt.Println("query 1::", queryDescontarStockOk)
		_, err = tx.ExecContext(ctx, queryDescontarStockOk)
		if err != nil {
			return !ok, err
		}
	}
	return ok, err
}

func (i *MySQLPedidoRepository) VerificarStockManufacturado(ctx context.Context, tx *sqlx.Tx, idArticulo, amount int) (bool, error) {
	var ok bool = true
	type VerificarStockInsumosQuery struct {
		Cantidad    int
		StockActual int
	}

	var verificarStockList []VerificarStockInsumosQuery

	queryArticuloManufacturado := `select ai.stock_actual, amd.cantidad from articulo_manufacturado am 
				JOIN articulo_manufacturado_detalle amd on amd.id_articulo_manufacturado = am.id
    			JOIN articulo_insumo ai on ai.id = amd.id_articulo_insumo
    			AND ai.es_insumo = true
    			AND am.id = ?`

	rows, err := tx.QueryxContext(ctx, queryArticuloManufacturado, idArticulo)
	if err != nil {
		return !ok, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock_actual, cantidad_pedida int
		if err := rows.Scan(&stock_actual, &cantidad_pedida); err != nil {
			return !ok, err
		}
		verificarStock := VerificarStockInsumosQuery{
			Cantidad:    cantidad_pedida,
			StockActual: stock_actual,
		}
		verificarStockList = append(verificarStockList, verificarStock)
	}

	for _, des := range verificarStockList {
		stockActual := strconv.Itoa(des.StockActual)
		cantidad := strconv.Itoa(des.Cantidad)
		idArticuloQ := strconv.Itoa(idArticulo)
		amountQ := strconv.Itoa(amount)

		//SELECT IF( ((stock_actual - ( amount * cantidad ) ) >= 0), true, false) from articulo_insumo WHERE id = 94;
		queryOk := []string{"SELECT IF( (( ", stockActual, " - (", amountQ,
			" * ", cantidad, ")) >= 0), true, false ) FROM articulo_insumo WHERE id = ", idArticuloQ}
		queryOkAux := strings.Join(queryOk, "")
		fmt.Println("query:", queryOkAux)
		rows, err := tx.QueryxContext(ctx, queryOkAux)
		if err != nil {
			return !ok, err
		}
		defer rows.Close()

		for rows.Next() {
			var ok bool
			if err := rows.Scan(&ok); err != nil {
				return ok, err
			}
			fmt.Println("Ok:", ok)
			if !ok {
				return ok, StockInsuficiente
			}
		}
	}
	return ok, err
}

func (i *MySQLPedidoRepository) CancelarPedido(ctx context.Context, tx *sqlx.Tx, idPedido int) error {
	queryCancelarPedido := "UPDATE pedidos SET estado = 6 WHERE id = ?"
	fmt.Println("queryCancelarPedido", queryCancelarPedido)
	_, err := tx.Exec(queryCancelarPedido, idPedido)
	if err != nil {
		return errors.New("no se logro cancelar el pedido correctamente")
	}
	fmt.Println("error", err)
	return err
}
