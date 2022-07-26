package domain

import "time"

type Factura struct {
	ID             int        `json:"id"`
	IDPedido       *int       `json:"id_pedido"`
	Fecha          *time.Time `json:"fecha"`
	NumeroFactura  *int       `json:"numero_factura"`
	MontoDescuento *float64   `json:"monto_descuento"`
	FormaPago      *string    `json:"forma_pago"`
	NumeroTarjeta  *string    `json:"numero_tarjeta"`
	TotalVenta     *float64   `json:"total_venta"`
	TotalCosto     *float64   `json:"total_costo"`
}
