package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type IFacturaService interface {
	GetAll(context.Context) ([]domain.Factura, error)
	GetByID(context.Context, int) (*domain.Factura, error)
	UpdateFactura(context.Context, domain.Factura) error
	DeleteFactura(context.Context, int) error
	AddFactura(context.Context, domain.Factura) error
	RecaudacionesDiarias(context.Context, string) ([]domain.Recaudaciones, error)
	RecaudacionesMensuales(context.Context, string, string) ([]domain.Recaudaciones, error)
	RecaudacionesPeriodoTiempo(context.Context, string, string) ([]domain.RecaudacionesResponse, error)
	ObtenerGanancias(context.Context, string, string) ([]domain.Ganancias, error)
}

type FacturaService struct {
	db         database.DB
	repository domain.IFacturaRepository
}

func NewFacturaService(db database.DB, repository domain.IFacturaRepository) *FacturaService {
	return &FacturaService{db, repository}
}

func (s *FacturaService) UpdateFactura(ctx context.Context, factura domain.Factura) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Update(ctx, tx, factura)
		return err
	})
	return err
}

func (s *FacturaService) GetByID(ctx context.Context, id int) (*domain.Factura, error) {
	var err error
	var factura *domain.Factura
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		factura, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("internal server error")
		}
		return err
	})
	return factura, err
}

func (s *FacturaService) AddFactura(ctx context.Context, factura domain.Factura) error {
	var err error
	fmt.Println("AddFactura service:", factura)
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, factura)
		return err
	})
	return err
}

func (s *FacturaService) GetAll(ctx context.Context) ([]domain.Factura, error) {
	var err error
	var facturas []domain.Factura
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		facturas, err = s.repository.GetAll(ctx, tx)
		return err
	})
	return facturas, err
}

func (s *FacturaService) RecaudacionesDiarias(ctx context.Context, fecha string) ([]domain.Recaudaciones, error) {
	var err error
	var recaudaciones []domain.Recaudaciones
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		recaudaciones, err = s.repository.RecaudacionesDiarias(ctx, tx, fecha)
		return err
	})
	return recaudaciones, err
}

func (s *FacturaService) RecaudacionesMensuales(ctx context.Context, month, year string) ([]domain.Recaudaciones, error) {
	var err error
	var recaudaciones []domain.Recaudaciones
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		recaudaciones, err = s.repository.RecaudacionesMensuales(ctx, tx, month, year)
		return err
	})
	return recaudaciones, err
}

func (s *FacturaService) RecaudacionesPeriodoTiempo(ctx context.Context, desde, hasta string) ([]domain.RecaudacionesResponse, error) {
	var err error
	var recaudacionesResponse []domain.RecaudacionesResponse
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		recaudacionesResponse, err = s.repository.RecaudacionesPeriodoTiempo(ctx, tx, desde, hasta)
		return err
	})
	return recaudacionesResponse, err
}

func (s *FacturaService) ObtenerGanancias(ctx context.Context, desde, hasta string) ([]domain.Ganancias, error) {
	var err error
	var ganancias []domain.Ganancias
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		ganancias, err = s.repository.ObtenerGanancias(ctx, tx, desde, hasta)
		return err
	})
	return ganancias, err
}

func (s *FacturaService) DeleteFactura(ctx context.Context, id int) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {

		_, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("factura not found")
		}
		err = s.repository.Delete(ctx, tx, id)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
