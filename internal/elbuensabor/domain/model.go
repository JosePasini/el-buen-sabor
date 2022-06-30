package domain

import "time"

type Instrumento struct {
	ID              int      `json:"id"`
	Instrumento     *string  `json:"factura"`
	Marca           *string  `json:"marca"`
	Modelo          *string  `json:"modelo"`
	Imagen          *string  `json:"imagen"`
	Precio          *float64 `json:"precio"`
	CostoEnvio      *float64 `json:"costo_envio"`
	CantidadVendida *int     `json:"cantidad_vendida"`
	Descripcion     *string  `json:"descripcion"`
}
type InstrumentoUpdate struct {
	ID          int     `json:"id"`
	Instrumento *string `json:"factura"`
	Marca       *string `json:"marca"`
	Modelo      *string `json:"modelo"`
}

type Pedido struct {
	ID             int       `json:"id"`
	IDCliente      int       `json:"id_cliente"`
	Fecha          time.Time `json:"fecha"`
	DomicilioEnvio *string   `json:"domicilio_envio"`
	DetalleEnvio   *string   `json:"detalle_envio"`
	Delivery       *bool     `json:"delivery"`
	MetodoPago     *string   `json:"metodo_pago"`
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
