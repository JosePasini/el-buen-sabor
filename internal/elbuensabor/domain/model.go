package domain

import "time"

type Pedido struct {
	ID              int       `json:"id"`
	Estado          int       `json:"estado"`
	HoraEstimadaFin time.Time `json:"hora_estimada_fin"`
	DetalleEnvio    *string   `json:"detalle_envio"`
	TipoEnvio       int       `json:"tipo_envio"`
	Total           float64   `json:"total"`
	IDDomicicio     int       `json:"id_domicilio"`
	IDCliente       int       `json:"id_cliente"`
}

type Factura struct {
	ID             int       `json:"id"`
	Fecha          time.Time `json:"fecha"`
	NumeroFactura  int       `json:"numero_factura"`
	MontoDescuento float64   `json:"monto_descuento"`
	FormaPago      *string   `json:"forma_pago"`
	NumeroTarjeta  *string   `json:"numero_tarjeta"`
	TotalVenta     float64   `json:"total_venta"`
	TotalCosto     float64   `json:"total_costo"`
}
