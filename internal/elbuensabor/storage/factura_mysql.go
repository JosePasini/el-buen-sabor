package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type recaudacionesDB struct {
	Recaudaciones sql.NullFloat64 `json:"recaudaciones" db:"recaudaciones"`
	Fecha         sql.NullString  `json:"fecha" db:"fecha"`
}

type recaudacionesResponseDB struct {
	Fecha         sql.NullString  `json:"fecha" db:"fecha"`
	NumeroFactura sql.NullInt32   `json:"numero_factura" db:"numero_factura"`
	FormaPago     sql.NullString  `json:"forma_pago" db:"forma_pago"`
	Recaudaciones sql.NullFloat64 `json:"recaudaciones" db:"recaudaciones"`
	IDPedido      sql.NullInt32   `json:"id_pedido" db:"id_pedido"`
}

func (i *recaudacionesResponseDB) toRecaudacionesResponse() domain.RecaudacionesResponse {
	return domain.RecaudacionesResponse{
		Fecha:         database.ToStringP(i.Fecha),
		NumeroFactura: database.ToIntP(i.NumeroFactura),
		FormaPago:     database.ToStringP(i.FormaPago),
		Recaudaciones: database.ToFloat64P(i.Recaudaciones),
		IDPedido:      database.ToIntP(i.IDPedido),
	}
}

type gananciasDB struct {
	Ganancias sql.NullFloat64 `json:"ganancias" db:"ganancias"`
	Desde     sql.NullString  `json:"desde" db:"desde"`
	Hasta     sql.NullString  `json:"hasta" db:"hasta"`
}

func (i *recaudacionesDB) toRecaudaciones() domain.Recaudaciones {
	return domain.Recaudaciones{
		Recaudaciones: database.ToFloat64P(i.Recaudaciones),
		Fecha:         database.ToStringP(i.Fecha),
	}
}

func (i *gananciasDB) toGanancias() domain.Ganancias {
	return domain.Ganancias{
		Ganancias: database.ToFloat64P(i.Ganancias),
		Desde:     database.ToStringP(i.Desde),
		Hasta:     database.ToStringP(i.Hasta),
	}
}

type facturaDB struct {
	ID             int             `db:"id"`
	Fecha          sql.NullTime    `db:"fecha"`
	NumeroFactura  sql.NullInt32   `db:"numero_factura"`
	MontoDescuento sql.NullFloat64 `db:"monto_descuento"`
	FormaPago      sql.NullString  `db:"forma_pago"`
	NumeroTarjeta  sql.NullString  `db:"numero_tarjeta"`
	TotalVenta     sql.NullFloat64 `db:"total_venta"`
	TotalCosto     sql.NullFloat64 `db:"total_costo"`
	IDPedido       sql.NullInt32   `db:"id_pedido"`
}

func (i *facturaDB) toFactura() domain.Factura {
	return domain.Factura{
		ID:             i.ID,
		Fecha:          database.ToTimeP(i.Fecha),
		NumeroFactura:  database.ToIntP(i.NumeroFactura),
		MontoDescuento: database.ToFloat64P(i.MontoDescuento),
		FormaPago:      database.ToStringP(i.FormaPago),
		NumeroTarjeta:  database.ToStringP(i.NumeroTarjeta),
		TotalVenta:     database.ToFloat64P(i.TotalVenta),
		TotalCosto:     database.ToFloat64P(i.TotalCosto),
		IDPedido:       database.ToIntP(i.IDPedido),
	}
}

type MySQLFacturaRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
	//qGetFacturaByIdPedido string
}

func NewMySQLFacturaRepository() *MySQLFacturaRepository {
	return &MySQLFacturaRepository{
		qInsert:     "INSERT INTO factura (fecha, numero_factura, monto_descuento, forma_pago, numero_tarjeta, total_venta, total_costo, id_pedido) VALUES (now(),?,?,?,?,?,?,?)",
		qGetByID:    "SELECT id, fecha, numero_factura, monto_descuento, forma_pago, numero_tarjeta, total_venta, total_costo FROM factura WHERE id = ?",
		qGetAll:     "SELECT id, fecha, numero_factura, monto_descuento, forma_pago, numero_tarjeta, total_venta, total_costo FROM factura",
		qDeleteById: "DELETE FROM factura WHERE id = ?",
		qUpdate:     "UPDATE factura SET fecha = COALESCE(?,fecha), numero_factura = COALESCE(?,numero_factura) , monto_descuento = COALESCE(?,monto_descuento), forma_pago = COALESCE(?,forma_pago), numero_tarjeta = COALESCE(?,numero_tarjeta), total_venta = COALESCE(?,total_venta), total_costo = COALESCE(?,total_costo) WHERE id = ?",
		//qGetFacturaByIdPedido: "SELECT id, fecha, numero_factura, monto_descuento, forma_pago, numero_tarjeta, total_venta, total_costo FROM factura WHERE id = ?",
	}
}

func (i *MySQLFacturaRepository) Update(ctx context.Context, tx *sqlx.Tx, fac domain.Factura) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, fac.Fecha, fac.NumeroFactura, fac.MontoDescuento, fac.FormaPago, fac.NumeroTarjeta, fac.TotalVenta, fac.TotalCosto, fac.ID)
	return err
}

func (i *MySQLFacturaRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLFacturaRepository) Insert(ctx context.Context, tx *sqlx.Tx, fac domain.Factura) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, fac.NumeroFactura, fac.MontoDescuento, fac.FormaPago, fac.NumeroTarjeta, fac.TotalVenta, fac.TotalCosto, fac.IDPedido)
	return err
}

func (i *MySQLFacturaRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Factura, error) {
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

func (i *MySQLFacturaRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Factura, error) {
	fmt.Println("Repository: 1")
	query := i.qGetAll
	facturas := make([]domain.Factura, 0)
	fmt.Println("Repository: 1")
	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println("Repository: 2")

	for rows.Next() {
		var factura facturaDB
		if err := rows.StructScan(&factura); err != nil {
			return facturas, err
		}
		facturas = append(facturas, factura.toFactura())
	}
	fmt.Println("Repository: 3")

	return facturas, nil
}

type facturaResponseDB struct {
	IDFactura  sql.NullInt32   `json:"id_factura" db:"id_factura"`
	Descuento  sql.NullFloat64 `json:"descuento" db:"descuento"`
	Fecha      sql.NullString  `json:"fecha" db:"fecha"`
	FormaPago  sql.NullString  `json:"forma_pago" db:"forma_pago"`
	TotalVenta sql.NullFloat64 `json:"total_venta" db:"total_venta"`
	Calle      sql.NullString  `json:"calle" db:"calle"`
	Numero     sql.NullInt32   `json:"numero" db:"numero"`
	Localidad  sql.NullString  `json:"localidad" db:"localidad"`
}

func (i *facturaResponseDB) toFacturaResponse() domain.FacturaResponse {
	return domain.FacturaResponse{
		IDFactura:  database.ToIntP(i.IDFactura),
		Fecha:      database.ToStringP(i.Fecha),
		FormaPago:  database.ToStringP(i.FormaPago),
		TotalVenta: database.ToFloat64P(i.TotalVenta),
		Descuento:  database.ToFloat64P(i.Descuento),
		Calle:      database.ToStringP(i.Calle),
		Numero:     database.ToIntP(i.Numero),
		Localidad:  database.ToStringP(i.Localidad),
	}
}

type pedidosResponseDB struct {
	Cantidad       sql.NullInt32   `json:"cantidad" db:"cantidad"`
	Denominacion   sql.NullString  `json:"denominacion" db:"denominacion"`
	PrecioUnitario sql.NullFloat64 `json:"precio" db:"precio"`
}

func (i *pedidosResponseDB) toPedidosResponse() domain.PedidoResponse {
	return domain.PedidoResponse{
		Cantidad:       database.ToIntP(i.Cantidad),
		Denominacion:   database.ToStringP(i.Denominacion),
		PrecioUnitario: database.ToFloat64P(i.PrecioUnitario),
	}
}

func (i *MySQLFacturaRepository) GetByIDPedido(ctx context.Context, tx *sqlx.Tx, idPedido int) (*domain.FacturaResponse, error) {
	query := `select f.id AS id_factura, f.fecha, f.forma_pago, f.total_venta, d.calle, d.numero, d.localidad from factura f 
		JOIN pedidos p ON p.id = f.id_pedido
		JOIN usuarios u ON u.id = p.id_cliente 
		JOIN domicilio d ON d.id_usuario = u.id
		WHERE f.id_pedido = ?;`
	var facturaResponse domain.FacturaResponse
	rows, err := tx.QueryxContext(ctx, query, idPedido)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var factura facturaResponseDB
		if err := rows.StructScan(&factura); err != nil {
			return &domain.FacturaResponse{}, err
		}
		facturaResponse = factura.toFacturaResponse()
	}

	var bebidas []domain.PedidoResponse
	var comidas []domain.PedidoResponse
	qGetBebidas := `SELECT dp.cantidad, am.denominacion, dp.subtotal AS precio FROM articulo_manufacturado am
			JOIN detalle_pedidos dp ON dp.id_articulo_manufacturado = am.id
			JOIN pedidos p ON p.id = dp.id_pedido
			WHERE p.id = ?`

	rows, err = tx.QueryxContext(ctx, qGetBebidas, idPedido)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pedido pedidosResponseDB
		if err := rows.StructScan(&pedido); err != nil {
			return &facturaResponse, err
		}
		bebidas = append(bebidas, pedido.toPedidosResponse())
	}
	if len(bebidas) > 0 {
		facturaResponse.Productos = append(facturaResponse.Productos, bebidas...)
	}

	qGetComidas := `SELECT dp.cantidad, ai.denominacion, dp.subtotal AS precio FROM articulo_insumo ai
			JOIN detalle_pedidos dp ON dp.id_articulo_insumo = ai.id
    		JOIN pedidos p ON p.id = dp.id_pedido
    		WHERE p.id = ?`
	rows, err = tx.QueryxContext(ctx, qGetComidas, idPedido)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pedido pedidosResponseDB
		if err := rows.StructScan(&pedido); err != nil {
			return &facturaResponse, err
		}
		comidas = append(comidas, pedido.toPedidosResponse())
	}
	if len(comidas) > 0 {
		facturaResponse.Productos = append(facturaResponse.Productos, comidas...)
	}
	fmt.Println("Facturas::::", facturaResponse)
	return &facturaResponse, nil
}

func (i *MySQLFacturaRepository) GetAllByCliente(ctx context.Context, tx *sqlx.Tx, idCliente int) ([]domain.FacturaResponse, error) {

	fmt.Println("Repository: 1")
	query := `select f.fecha, f.forma_pago, f.total_venta, d.calle, d.numero, d.localidad from factura f 
		JOIN pedidos p ON p.id = f.id_pedido
		JOIN usuarios u ON u.id = p.id_cliente 
		JOIN domicilio d ON d.id_usuario = u.id
		WHERE f.id_pedido = ?;`

	facturas := make([]domain.FacturaResponse, 0)
	fmt.Println("Repository: 1")
	rows, err := tx.QueryxContext(ctx, query, idCliente)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println("Repository: 2")

	for rows.Next() {
		var factura facturaResponseDB
		if err := rows.StructScan(&factura); err != nil {
			return facturas, err
		}
		facturas = append(facturas, factura.toFacturaResponse())
	}
	fmt.Println("Repository: 3")
	fmt.Println("Facturas::::", facturas)
	return facturas, nil
}

func (i *MySQLFacturaRepository) RecaudacionesDiarias(ctx context.Context, tx *sqlx.Tx, fecha string) ([]domain.Recaudaciones, error) {
	query_recaudacion_diaria := []string{"SELECT sum(total_venta) AS recaudaciones, fecha FROM factura WHERE fecha LIKE '", fecha, "%';"}
	qRecaudacionDiaria := strings.Join(query_recaudacion_diaria, "")
	recaudaciones := make([]domain.Recaudaciones, 0)

	fmt.Println("query:", qRecaudacionDiaria)
	fmt.Println("fecha:", fecha)
	rows, err := tx.QueryxContext(ctx, qRecaudacionDiaria)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rec recaudacionesDB
		if err := rows.StructScan(&rec); err != nil {
			return recaudaciones, err
		}
		reca := rec.toRecaudaciones()
		recaudaciones = append(recaudaciones, reca)
	}
	return recaudaciones, nil
}

func (i *MySQLFacturaRepository) RecaudacionesMensuales(ctx context.Context, tx *sqlx.Tx, month, year string) ([]domain.Recaudaciones, error) {
	query_recaudacion_mensual := []string{"SELECT SUM(total_venta), fecha FROM factura WHERE EXTRACT(MONTH FROM factura.fecha) =", month, " AND EXTRACT(YEAR FROM factura.fecha) =", year}
	qRecaudacionMensual := strings.Join(query_recaudacion_mensual, "")
	recaudaciones := make([]domain.Recaudaciones, 0)

	fmt.Println("query:", qRecaudacionMensual)
	fmt.Println("mes:", month)
	fmt.Println("año:", year)
	rows, err := tx.QueryxContext(ctx, qRecaudacionMensual)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var reca recaudacionesDB
		if err := rows.StructScan(&reca); err != nil {
			return recaudaciones, err
		}

		newRecaudacion := reca.toRecaudaciones()
		recaudaciones = append(recaudaciones, newRecaudacion)
	}
	return recaudaciones, nil
}

func (i *MySQLFacturaRepository) RecaudacionesPeriodoTiempo(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]domain.RecaudacionesResponse, error) {
	qRecaudacionPeriodoTiempo := `select fecha, numero_factura, forma_pago, total_venta AS recaudaciones, id_pedido from factura 
			where fecha BETWEEN ? AND ?`
	recaudaciones := make([]domain.RecaudacionesResponse, 0)

	fmt.Println("query:", qRecaudacionPeriodoTiempo)
	fmt.Println("desde:", desde)
	fmt.Println("hasta:", hasta)
	rows, err := tx.QueryxContext(ctx, qRecaudacionPeriodoTiempo, desde, hasta)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recau recaudacionesResponseDB
		if err := rows.StructScan(&recau); err != nil {
			return recaudaciones, err
		}

		newRecaudacion := recau.toRecaudacionesResponse()
		recaudaciones = append(recaudaciones, newRecaudacion)
	}
	return recaudaciones, nil
}

func (i *MySQLFacturaRepository) ObtenerGanancias(ctx context.Context, tx *sqlx.Tx, desde, hasta string) ([]domain.Ganancias, error) {
	//query_ganancia := []string{"SELECT SUM(total_venta) - sum(total_costo) AS ganancias FROM factura WHERE fecha BETWEEN", month, " AND EXTRACT(YEAR FROM factura.fecha) =", year}
	//qGanancias := strings.Join(query_ganancia, "")
	qGanancias := "SELECT SUM(total_venta) - sum(total_costo) AS ganancias FROM factura WHERE fecha BETWEEN ? AND ?"

	gananciasResponse := make([]domain.Ganancias, 0)
	var ganancias domain.Ganancias
	fmt.Println("query:", qGanancias)
	fmt.Println("desde:", desde)
	fmt.Println("hasta:", hasta)
	rows, err := tx.QueryxContext(ctx, qGanancias, desde, hasta)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ganancia gananciasDB
		if err := rows.StructScan(&ganancia); err != nil {
			return nil, err
		}
		ganancias = ganancia.toGanancias()
		ganancias.Desde = &desde
		ganancias.Hasta = &hasta
		// ganancias = domain.Ganancias{
		// 	Ganancias: gan,
		// 	Desde:     desde,
		// 	Hasta:     hasta,
		// }
		gananciasResponse = append(gananciasResponse, ganancias)
	}

	return gananciasResponse, nil
}
