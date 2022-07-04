package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
	"github.com/jmoiron/sqlx"
)

type IArticuloInsumoService interface {
	GetAll(context.Context) ([]domain.ArticuloInsumo, error)
	GetByID(context.Context, int) (*domain.ArticuloInsumo, error)
	UpdateArticuloInsumo(context.Context, domain.ArticuloInsumo) error
	DeleteArticuloInsumo(context.Context, int) error
	AddArticuloInsumo(context.Context, domain.ArticuloInsumo) error
}

type ArticuloInsumoService struct {
	db         database.DB
	repository storage.IArticuloInsumoRepository
}

func NewArticuloInsumoService(db database.DB, repository storage.IArticuloInsumoRepository) *ArticuloInsumoService {
	return &ArticuloInsumoService{db, repository}
}

func (s *ArticuloInsumoService) UpdateArticuloInsumo(ctx context.Context, articuloInsumo domain.ArticuloInsumo) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Update(ctx, tx, articuloInsumo)
		return err
	})
	return err
}

func (s *ArticuloInsumoService) GetByID(ctx context.Context, id int) (*domain.ArticuloInsumo, error) {
	var err error
	var articuloInsumo *domain.ArticuloInsumo
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		articuloInsumo, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("internal server error")
		}
		return err
	})
	return articuloInsumo, err
}

func (s *ArticuloInsumoService) AddArticuloInsumo(ctx context.Context, articuloInsumo domain.ArticuloInsumo) error {
	var err error
	fmt.Println("AddArticuloInsumo service:", articuloInsumo)
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, articuloInsumo)
		return err
	})
	return err
}

func (s *ArticuloInsumoService) GetAll(ctx context.Context) ([]domain.ArticuloInsumo, error) {
	var err error
	var articuloInsumo []domain.ArticuloInsumo
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		articuloInsumo, err = s.repository.GetAll(ctx, tx)
		return err
	})
	return articuloInsumo, err
}

func (s *ArticuloInsumoService) DeleteArticuloInsumo(ctx context.Context, id int) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {

		_, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("articulo not found")
		}
		err = s.repository.Delete(ctx, tx, id)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
