package domain

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
