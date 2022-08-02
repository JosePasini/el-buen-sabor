package domain

import "time"

type UsuarioResponse struct {
	ID       int     `json:"id"`
	Nombre   *string `json:"nombre"`
	Apellido *string `json:"apellido"`
	Usuario  *string `json:"usuario"`
	Email    *string `json:"email"`
	Rol      int     `json:"rol"`
}

type RankingComidasMasPedidas struct {
	IDPedido                int    `json:"id_pedido" db:"id_pedido"`
	VecesPedida             int    `json:"veces_pedida" db:"veces_pedida"`
	IDArticuloManufacturado int    `json:"id_articulo_manufacturado" db:"id_articulo_manufacturado"`
	Denominacion            string `json:"denominacion" db:"denominacion"`
}

type DetallePedidoResponse struct {
	IDPedido     int     `json:"id_pedido" db:"id_pedido"`
	Cantidad     int     `json:"cantidad" db:"cantidad"`
	Subtotal     float64 `json:"subtotal" db:"subtotal"`
	Denominacion *string `json:"denominacion" db:"denominacion"`
	Imagen       *string `json:"imagen" db:"imagen"`
}

type PedidosPorCliente struct {
	CantidadPedidos int     `json:"cantidad_pedidos" db:"cantidad_pedidos"`
	IDCliente       int     `json:"id_cliente" db:"id_cliente"`
	Total           float64 `json:"total" db:"total"`
}

type Recaudaciones struct {
	Recaudaciones *float64 `json:"recaudaciones" db:"recaudaciones"`
	Fecha         *string  `json:"fecha" db:"fecha"`
}

type Ganancias struct {
	Ganancias *float64 `json:"ganancias" db:"ganancias"`
	Desde     *string  `json:"desde" db:"desde"`
	Hasta     *string  `json:"hasta" db:"hasta"`
}

type RecaudacionesResponse struct {
	Fecha         *string  `json:"fecha" db:"fecha"`
	NumeroFactura *int     `json:"numero_factura" db:"numero_factura"`
	FormaPago     *string  `json:"forma_pago" db:"forma_pago"`
	Recaudaciones *float64 `json:"recaudaciones" db:"recaudaciones"`
	IDPedido      *int     `json:"id_pedido" db:"id_pedido"`
}

type FacturaResponse struct {
	//FacturaAuxResponse FacturaAuxResponse `json:"factura"`
	IDFactura  *int             `json:"id_factura" db:"id_factura"`
	Descuento  *float64         `json:"descuentos" db:"descuentos"`
	Fecha      *string          `json:"fecha" db:"fecha"`
	FormaPago  *string          `json:"forma_pago" db:"forma_pago"`
	TotalVenta *float64         `json:"total_venta" db:"total_venta"`
	Calle      *string          `json:"calle" db:"calle"`
	Numero     *int             `json:"numero" db:"numero"`
	Localidad  *string          `json:"localidad" db:"localidad"`
	Productos  []PedidoResponse `json:"productos"`
}

type FacturaAuxResponse struct {
	Domicilio *string   `json:"domicilio" db:"domicilio"`
	Pago      *string   `json:"pago" db:"pago"`
	Fecha     time.Time `json:"fecha" db:"fecha"`
	Total     *float64  `json:"total" db:"total"`
}

type PedidoResponse struct {
	Cantidad       *int     `json:"cantidad" db:"cantidad"`
	Denominacion   *string  `json:"denominacion" db:"denominacion"`
	PrecioUnitario *float64 `json:"precio" db:"precio"`
}
