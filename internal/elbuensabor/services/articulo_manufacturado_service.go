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
	GetAllAvailable(context.Context) (map[string][]domain.ArticuloManufacturadoAvailable, error)
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

func (i *ArticuloManufacturadoService) GetAllAvailable(ctx context.Context) (map[string][]domain.ArticuloManufacturadoAvailable, error) {
	var err error
	var bandera bool = false
	var articulosManufacturadosAvailable []*domain.ArticuloManufacturadoAvailable
	articulosMap := make(map[string][]domain.ArticuloManufacturadoAvailable)

	err = i.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		articulosManufacturadosAvailable, err = i.repository.GetAllAvailable(ctx, tx)
		return err
	})

	for _, articulo := range articulosManufacturadosAvailable {
		if *articulo.CantidadNecesaria > *articulo.StockActual {
			articulo.Disponible = &bandera
		}
		key, exists := articulosMap[*articulo.ArticuloManufacturado]
		if exists {
			key = append(key, *articulo)
		} else {
			key = []domain.ArticuloManufacturadoAvailable{*articulo}
		}
		articulosMap[*articulo.ArticuloManufacturado] = key
	}

	return articulosMap, err
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
