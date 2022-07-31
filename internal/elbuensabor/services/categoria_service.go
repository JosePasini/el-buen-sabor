package services

import (
	"context"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
	"github.com/jmoiron/sqlx"
)

type ICategoriaService interface {
	AddCategoria(context.Context, domain.Categoria) error
	GetAll(context.Context) ([]domain.Categoria, error)
}

type CategoriaService struct {
	db         database.DB
	repository storage.ICategoriaRepository
}

func NewCategoriaService(db database.DB, repository storage.ICategoriaRepository) *CategoriaService {
	return &CategoriaService{db, repository}
}

func (s *CategoriaService) AddCategoria(ctx context.Context, categoria domain.Categoria) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, categoria)
		return err
	})
	return err
}

func (s *CategoriaService) GetAll(ctx context.Context) ([]domain.Categoria, error) {
	var err error
	var categorias []domain.Categoria
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		categorias, err = s.repository.GetAll(ctx, tx)
		return err
	})
	return categorias, err
}
