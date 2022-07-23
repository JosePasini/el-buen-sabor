package domain

type Domicilio struct {
	ID        int    `json:"id" db:"id"`
	Calle     string `json:"calle" db:"calle"`
	Numero    int    `json:"numero" db:"numero"`
	Localidad string `json:"localidad" db:"localidad"`
}
