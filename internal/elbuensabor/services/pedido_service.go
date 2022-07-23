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

type IPedidoService interface {
	GetAll(context.Context) ([]domain.Pedido, error)
	GetByID(context.Context, int) (*domain.Pedido, error)
	UpdatePedido(context.Context, domain.Pedido) error
	DeletePedido(context.Context, int) error
	AddPedido(context.Context, domain.Pedido) error
	GenerarPedido(context.Context, domain.GenerarPedido) (domain.GenerarPedido, error)
	AceptarPedido(context.Context, int) (bool, error)
	RankingComidasMasPedidas(context.Context) ([]domain.RankingComidasMasPedidas, error)
}

type PedidoService struct {
	db         database.DB
	repository storage.IPedidoRepository
}

func NewPedidoService(db database.DB, repository storage.IPedidoRepository) *PedidoService {
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
		_, err = s.repository.Insert(ctx, tx, pedido)
		return err
	})
	return nil
}

func (s *PedidoService) GenerarPedido(ctx context.Context, generarPedido domain.GenerarPedido) (domain.GenerarPedido, error) {
	var err error
	var pedido = generarPedido.Pedido
	var detallePedido = generarPedido.DetallePedido
	var idPedido int
	var total int
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		// creamos un 'pedido' en la BD y nos retorna el ID
		idPedido, err = s.repository.Insert(ctx, tx, pedido)

		// insertamos todos los detalles con el ID de pedido.
		if len(detallePedido) > 0 {
			for _, detalle := range detallePedido {
				detalle.IDPedido = idPedido
				err = s.repository.InsertDetallePedido(ctx, tx, detalle)
				total += int(detalle.SubTotal) * detalle.Cantidad
				if err != nil {
					return err
				}
			}
		}

		// updateamos el pedido con el total final.
		err = s.repository.UpdateTotal(ctx, tx, total, idPedido)
		return err
	})
	if err != nil {
		return domain.GenerarPedido{}, err
	}
	return generarPedido, err
}

func (s *PedidoService) AceptarPedido(ctx context.Context, idPedido int) (bool, error) {
	var err error
	var pedido *domain.Pedido
	var ok bool = true
	fmt.Println("Aceptar pedido service")
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		// obtengo el pedido por id para comprobar que exista
		pedido, err = s.repository.GetByID(ctx, tx, idPedido)
		if err != nil {
			return errors.New("internal server error")
		}
		// comprueba que el estado del pedido sea 1 :: 'pendiente de aprobacion'
		if pedido.Estado != 1 {
			return errors.New("solo se puede aceptar pedidos en estado 'pendiente de aprobacion' :: 1 ")
		}

		ok, err = s.repository.DescontarStock(ctx, tx, idPedido)
		if err != nil || !ok {
			return err
		}
		return err
	})
	fmt.Println(err)
	if err != nil {
		return false, err
	}
	fmt.Println("pedido", pedido)
	return true, nil
}

func (s *PedidoService) RankingComidasMasPedidas(ctx context.Context) ([]domain.RankingComidasMasPedidas, error) {
	var err error
	var rankingComidasMasPedidas []domain.RankingComidasMasPedidas
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		rankingComidasMasPedidas, err = s.repository.RankingComidasMasPedidas(ctx, tx)
		return err
	})
	return rankingComidasMasPedidas, err
}
