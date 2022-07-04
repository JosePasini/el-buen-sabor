package domain

type ArticuloInsumo struct {
	ID           int     `json:"id"`
	Denominacion *string `json:"denominacion"`
}

type ArticuloManufacturadoDetalle struct {
	ID int `json:"id"`
}

type ArticuloManufacturado struct {
	ID int `json:"id"`
}
