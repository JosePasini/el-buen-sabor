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
