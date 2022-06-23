package domain

import "time"

type Instrumento struct {
	ID              int      `json:"id"`
	Instrumento     *string  `json:"instrumento"`
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
	Instrumento *string `json:"instrumento"`
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
