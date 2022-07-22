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
	ID                   int      `json:"id"`
	TiempoEstimadoCocina *int     `json:"tiempo_estimado_cocina"`
	Denominacion         *string  `json:"denominacion"`
	PrecioVenta          *float64 `json:"precio_venta"`
	Imagen               *string  `json:"imagen"`
}

type ArticuloManufacturadoAvailable struct {
	CantidadNecesaria     *int     `json:"cantidad_necesaria"`
	UnidadMedida          *string  `json:"unidad_medida"`
	Insumo                *string  `json:"insumo"`
	StockActual           *int     `json:"stock_actual"`
	ArticuloManufacturado *string  `json:"articulo_manufacturado"`
	TiempoEstimadoCocina  *int     `json:"tiempo_estimado_cocina"`
	PrecioVenta           *float64 `json:"precio_venta"`
	Disponible            *bool    `json:"disponible" default:"true"`
}

/*cantidad: null
id: null
id_articulo_insumo: null
id_articulo_manufacturado: null
subtotal: null
*/
// select id, denominacion, precio_compra, precio_venta, stock_actual, stock_minimo, unidad_medida, es_insumo, imagen from articulo_insumo WHERE es_insumo = false;
// select id, tiempo_estimado_cocina, denominacion, precio_venta, imagen from articulo_manufacturado;
type CarritoCompleto struct {
	ID                   int
	Denominacion         *string
	PrecioCompra         *float64
	PrecioVenta          *float64
	Cantidad             int
	StockActual          *int
	StockMinimo          *int
	Imagen               *string
	EsBebida             bool
	TiempoEstimadoCocina *int
}
