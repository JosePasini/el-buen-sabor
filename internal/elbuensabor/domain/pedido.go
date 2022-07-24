package domain

import "time"

const (
	PENDIENTE_APROBACION  int = 1
	APROBADO_COCINA       int = 2
	PENDIENTE_DE_DESPACHO int = 3
	DELIVERY_EN_CAMINO    int = 4
	FACTURADO             int = 5
	DENEGADO              int = 6
)

const (
	DELIVERY     int = 1
	RETIRO_LOCAL int = 2
)

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
