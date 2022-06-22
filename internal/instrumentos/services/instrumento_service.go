package services

import (
	"context"
	"errors"

	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/database"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/domain"
	"github.com/jmoiron/sqlx"
)

type IInstrumentoService interface {
	GetAll(context.Context) ([]domain.Instrumento, error)
	GetByID(context.Context, int) (*domain.Instrumento, error)
	UpdateInstrument(context.Context, domain.Instrumento) error
	DeleteInstrument(context.Context, int) error
	AddInstrumento(context.Context, domain.Instrumento) error
}

type InstrumentoService struct {
	db         database.DB
	repository domain.IInstrumentoRepository
}

func NewInstrumentoService(db database.DB, repository domain.IInstrumentoRepository) *InstrumentoService {
	return &InstrumentoService{db, repository}
}

func (s *InstrumentoService) UpdateInstrument(ctx context.Context, instrumento domain.Instrumento) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Update(ctx, tx, instrumento)
		return err
	})
	return err
}

func (s *InstrumentoService) GetByID(ctx context.Context, id int) (*domain.Instrumento, error) {
	var err error
	var instrumento *domain.Instrumento
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		instrumento, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("internal server error")
		}
		return err
	})
	return instrumento, err
}

func (s *InstrumentoService) AddInstrumento(ctx context.Context, instrumento domain.Instrumento) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, instrumento)
		return err
	})
	return err
}

func (s *InstrumentoService) GetAll(ctx context.Context) ([]domain.Instrumento, error) {
	var err error
	var instruments []domain.Instrumento
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		instruments, err = s.repository.GetAll(ctx, tx)
		return err
	})
	return instruments, err
}

func (s *InstrumentoService) DeleteInstrument(ctx context.Context, id int) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {

		_, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("instrument not found")
		}
		err = s.repository.Delete(ctx, tx, id)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
