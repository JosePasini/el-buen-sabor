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

type IArticuloManufacturadoService interface {
	GetAll(context.Context) ([]domain.ArticuloManufacturado, error)
	GetByID(context.Context, int) (*domain.ArticuloManufacturado, error)
	UpdateArticuloManufacturado(context.Context, domain.ArticuloManufacturado) error
	DeleteArticuloManufacturado(context.Context, int) error
	AddArticuloManufacturado(context.Context, domain.ArticuloManufacturado) error
}

type ArticuloManufacturadoService struct {
	db         database.DB
	repository storage.IArticuloManufacturadoRepository
}

func NewArticuloManufacturadoService(db database.DB, repository storage.IArticuloManufacturadoRepository) *ArticuloManufacturadoService {
	return &ArticuloManufacturadoService{db, repository}
}

func (i *ArticuloManufacturadoService) GetAll(ctx context.Context) ([]domain.ArticuloManufacturado, error) {
	var err error
	var articulosManufacturados []domain.ArticuloManufacturado
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		articulosManufacturados, err = i.repository.GetAll(ctx, tx)
		return err
	})
	return articulosManufacturados, err
}

func (i *ArticuloManufacturadoService) GetByID(ctx context.Context, id int) (*domain.ArticuloManufacturado, error) {
	var err error
	var articulosManufacturados *domain.ArticuloManufacturado
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		articulosManufacturados, err = i.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("internal server error")
		}
		return err
	})
	return articulosManufacturados, err
}

func (i *ArticuloManufacturadoService) UpdateArticuloManufacturado(ctx context.Context, articuloManufacturado domain.ArticuloManufacturado) error {
	var err error
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = i.repository.Update(ctx, tx, articuloManufacturado)
		return err
	})
	return err
}

func (i *ArticuloManufacturadoService) DeleteArticuloManufacturado(ctx context.Context, id int) error {
	var err error
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {

		_, err = i.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("factura not found")
		}
		err = i.repository.Delete(ctx, tx, id)
		if err != nil {
			return err
		}
		return err
	})
	return err
}

func (i *ArticuloManufacturadoService) AddArticuloManufacturado(ctx context.Context, articuloManufacturado domain.ArticuloManufacturado) error {
	var err error
	fmt.Println("articuloManufacturado service:", articuloManufacturado)
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = i.repository.Insert(ctx, tx, articuloManufacturado)
		return err
	})
	return err
}
