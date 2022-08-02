package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type categoriaDB struct {
	ID       int            `json:"id" db:"id"`
	Nombre   sql.NullString `json:"nombre" db:"nombre"`
	EsInsumo sql.NullBool   `json:"es_insumo" db:"es_insumo"`
}

func (a *categoriaDB) toCategoria() domain.Categoria {
	return domain.Categoria{
		ID:       a.ID,
		Nombre:   database.ToStringP(a.Nombre),
		EsInsumo: database.ToBoolP(a.EsInsumo),
	}
}

type ICategoriaRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, categoria domain.Categoria) error
	GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Categoria, error)
}

type MySQLCategoriaRepository struct {
	qInsert string
	qGetAll string
}

func NewMySQLCategoriaRepository() *MySQLCategoriaRepository {
	return &MySQLCategoriaRepository{
		qInsert: "INSERT INTO categoria (nombre, es_insumo) VALUES (?,?)",
		qGetAll: "SELECT id, nombre, es_insumo FROM categoria",
	}
}

func (i *MySQLCategoriaRepository) Insert(ctx context.Context, tx *sqlx.Tx, categoria domain.Categoria) error {
	fmt.Println("domicilio:", categoria)
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, categoria.Nombre, categoria.EsInsumo)
	return err
}

func (i *MySQLCategoriaRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Categoria, error) {
	query := i.qGetAll
	categorias := make([]domain.Categoria, 0)

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat categoriaDB
		if err := rows.StructScan(&cat); err != nil {
			return categorias, err
		}
		categorias = append(categorias, cat.toCategoria())
	}
	return categorias, nil

}
