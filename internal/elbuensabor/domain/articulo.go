package domain

type ArticuloInsumo struct {
	ID           int     `json:"id"`
	Denominacion *string `json:"denominacion"`
}

type ArticuloManufacturadoDetalle struct {
	ID           int    `json:"id"`
	UnidadMedida string `json:"unidad_medida"`
}

type ArticuloManufacturado struct {
	ID           int     `json:"id"`
	Denominacion *string `json:"denominacion"`
}
