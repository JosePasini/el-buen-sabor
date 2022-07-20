package domain

type Login struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type ProductoMercadoPago struct {
	ID                   int     `json:"id"`
	PrecioVenta          float64 `json:"precio_venta"`
	TiempoEstimadoCocina int     `json:"tiempo_estimado_cocina"`
	Amount               int     `json:"amount"`
	Imagen               string  `json:"imagen"`
	Denominacion         string  `json:"denominacion"`
}

type UsuarioMercadoPago struct {
	ID                   int    `json:"id"`
	Nombre               string `json:"nombre"`
	Apellido             string `json:"apellido"`
	Usuario              string `json:"usuario"`
	Email                string `json:"email"`
	Rol                  int    `json:"rol"`
	TypeIdentification   string `json:"type_identification"`
	NumberIdentification string `json:"number_identification"`
}

type DetallePedido struct {
	ID                      int     `json:"id"`
	Cantidad                int     `json:"cantidad"`
	Subtotal                float64 `json:"subtotal"`
	IdArticuloManufacturado int     `json:"id_articulo_manufacturado"`
	IdArticuloInsumo        int     `json:"id_articulo_insumo"`
	IdPedido                int     `json:"id_pedido"`
}

type GenerarPedido struct {
	Pedido        Pedido          `json:"pedido"`
	DetallePedido []DetallePedido `json:"detalle_pedido"`
}
