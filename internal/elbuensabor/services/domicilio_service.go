package services

import (
	"context"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
	"github.com/jmoiron/sqlx"
)

type IDomicilioService interface {
	AddDomicilio(context.Context, domain.Domicilio) error
	UpdateDomicilio(context.Context, domain.Domicilio) error
}

type DomicilioService struct {
	db         database.DB
	repository storage.IDomicilioRepository
}

func NewDomicilioService(db database.DB, repository storage.IDomicilioRepository) *DomicilioService {
	return &DomicilioService{db, repository}
}

func (s *DomicilioService) AddDomicilio(ctx context.Context, domicilio domain.Domicilio) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, domicilio)
		return err
	})
	return err
}

func (s *DomicilioService) UpdateDomicilio(ctx context.Context, domicilio domain.Domicilio) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Update(ctx, tx, domicilio)
		return err
	})
	return err
}
