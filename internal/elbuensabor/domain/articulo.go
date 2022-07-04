package domain

type ArticuloInsumo struct {
	ID           int      `json:"id"`
	Denominacion *string  `json:"denominacion"`
	PrecioCompra *float64 `json:"precio_compra"`
	PrecioVenta  *float64 `json:"precio_venta"`
	StockActual  *int     `json:"stock_actual"`
	StockMinimo  *int     `json:"stock_minimo"`
	UnidadMedida *string  `json:"unidad_medida"`
	EsInsumo     *bool    `json:"es_insumo"`
}

type ArticuloManufacturadoDetalle struct {
	ID                      int `json:"id"`
	Cantidad                int `json:"cantidad"`
	IDArticuloManufacturado int `json:"id_articulo_manufacturado"`
	IDArticuloInsumo        int `json:"id_articulo_insumo"`
}

type ArticuloManufacturado struct {
	ID           int      `json:"id"`
	Denominacion *string  `json:"denominacion"`
	PrecioVenta  *float64 `json:"precio_venta"`
	Imagen       *string  `json:"imagen"`
}
