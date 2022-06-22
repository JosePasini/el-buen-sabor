package domain

type Instrumento struct {
	Id              int      `json:"id"`
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
	Id          int     `json:"id"`
	Instrumento *string `json:"instrumento"`
	Marca       *string `json:"marca"`
	Modelo      *string `json:"modelo"`
}
