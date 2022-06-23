package services

import (
	"context"
	"errors"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type IPedidoService interface {
	GetAll(context.Context) ([]domain.Pedido, error)
	GetByID(context.Context, int) (*domain.Pedido, error)
	UpdatePedido(context.Context, domain.Pedido) error
	DeletePedido(context.Context, int) error
	AddPedido(context.Context, domain.Pedido) error
}

type PedidoService struct {
	db         database.DB
	repository domain.IPedidoRepository
}

func NewPedidoService(db database.DB, repository domain.IPedidoRepository) *PedidoService {
	return &PedidoService{db, repository}
}

func (s *PedidoService) GetAll(ctx context.Context) ([]domain.Pedido, error) {
	var err error
	var pedidos []domain.Pedido
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		pedidos, err = s.repository.GetAll(ctx, tx)
		return err
	})
	return pedidos, err
}

func (s *PedidoService) GetByID(ctx context.Context, id int) (*domain.Pedido, error) {
	var err error
	var pedido *domain.Pedido
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		pedido, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("internal server error")
		}
		return err
	})
	return pedido, err
}

func (s *PedidoService) UpdatePedido(ctx context.Context, pedido domain.Pedido) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Update(ctx, tx, pedido)
		return err
	})
	return err
}

func (s *PedidoService) DeletePedido(ctx context.Context, id int) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		_, err = s.repository.GetByID(ctx, tx, id)
		if err != nil {
			return errors.New("pedido not found")
		}
		err = s.repository.Delete(ctx, tx, id)
		if err != nil {
			return err
		}
		return err
	})
	return err
}

func (s *PedidoService) AddPedido(ctx context.Context, pedido domain.Pedido) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, pedido)
		return err
	})
	return nil
}
