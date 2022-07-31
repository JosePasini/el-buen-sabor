package domain

type Categoria struct {
	ID     int     `json:"id" bd:"id"`
	Nombre *string `json:"nombre" bd:"nombre"`
}
