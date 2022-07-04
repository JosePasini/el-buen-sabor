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

type IArticuloManufacturadoDetalleService interface {
	GetAll(ctx context.Context) ([]domain.ArticuloManufacturadoDetalle, error)
	GetByID(ctx context.Context, id int) (*domain.ArticuloManufacturadoDetalle, error)
	UpdateArticuloManufacturadoDetalle(ctx context.Context, artManufacturado domain.ArticuloManufacturadoDetalle) error
	DeleteArticuloManufacturadoDetalle(ctx context.Context, id int) error
	AddArticuloManufacturadoDetalle(ctx context.Context, artManufacturado domain.ArticuloManufacturadoDetalle) error
}

type ArticuloManufacturadoDetalleService struct {
	db         database.DB
	repository storage.IArticuloManufacturadoDetalleRepository
}

func NewArticuloManufacturadoDetalleService(db database.DB, repository storage.IArticuloManufacturadoDetalleRepository) *ArticuloManufacturadoDetalleService {
	return &ArticuloManufacturadoDetalleService{db, repository}
}

func (i *ArticuloManufacturadoDetalleService) GetAll(ctx context.Context) ([]domain.ArticuloManufacturadoDetalle, error) {
	var err error
	var articulosManufacturados []domain.ArticuloManufacturadoDetalle
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		articulosManufacturados, err = i.repository.GetAll(ctx, tx)
		return err
	})
	return articulosManufacturados, err
}

func (i *ArticuloManufacturadoDetalleService) GetByID(ctx context.Context, id int) (*domain.ArticuloManufacturadoDetalle, error) {
	var err error
	var articulosManufacturados *domain.ArticuloManufacturadoDetalle
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		articulosManufacturados, err = i.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("internal server error")
		}
		return err
	})
	return articulosManufacturados, err
}

func (i *ArticuloManufacturadoDetalleService) UpdateArticuloManufacturadoDetalle(ctx context.Context, articuloManufacturado domain.ArticuloManufacturadoDetalle) error {
	var err error
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = i.repository.Update(ctx, tx, articuloManufacturado)
		return err
	})
	return err
}

func (i *ArticuloManufacturadoDetalleService) DeleteArticuloManufacturadoDetalle(ctx context.Context, id int) error {
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

func (i *ArticuloManufacturadoDetalleService) AddArticuloManufacturadoDetalle(ctx context.Context, articuloManufacturado domain.ArticuloManufacturadoDetalle) error {
	var err error
	fmt.Println("articuloManufacturado service:", articuloManufacturado)
	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = i.repository.Insert(ctx, tx, articuloManufacturado)
		return err
	})
	return err
}
