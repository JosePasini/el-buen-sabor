package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
	"github.com/jmoiron/sqlx"
)

type IPedidoService interface {
	GetAll(context.Context) ([]domain.Pedido, error)
	GetAllPedidosByIDCliente(context.Context, int) ([]domain.Pedido, error)
	GetByID(context.Context, int) (*domain.Pedido, error)
	UpdatePedido(context.Context, domain.Pedido) error
	UpdateEstadoPedido(context.Context, int, int) error
	DeletePedido(context.Context, int) error
	//AddPedido(context.Context, domain.Pedido) error
	GenerarPedido(context.Context, domain.GenerarPedido) (domain.GenerarPedido, error)
	AceptarPedido(context.Context, int) (bool, error)
	VerificarStock(context.Context, int, int, bool) (bool, error)
	CancelarPedido(context.Context, int) error
	RankingComidasMasPedidas(context.Context, string, string) ([]domain.RankingComidasMasPedidas, error)
	GetAllDetallePedidosByIDPedido(context.Context, int) ([]domain.DetallePedidoResponse, error)
	GetRankingDePedidosPorCliente(context.Context, string, string) ([]domain.PedidosPorCliente, error)
}

type PedidoService struct {
	db                database.DB
	repository        storage.IPedidoRepository
	repositoryFactura domain.IFacturaRepository
	repositoryLogin   domain.ILoginRepository
}

func NewPedidoService(db database.DB, repository storage.IPedidoRepository, repositoryFactura domain.IFacturaRepository, repositoryLogin domain.ILoginRepository) *PedidoService {
	return &PedidoService{db, repository, repositoryFactura, repositoryLogin}
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
func (s *PedidoService) GetAllPedidosByIDCliente(ctx context.Context, idCliente int) ([]domain.Pedido, error) {
	var err error
	var pedidos []domain.Pedido
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		pedidos, err = s.repository.GetAllPedidosByIDCliente(ctx, tx, idCliente)
		return err
	})
	return pedidos, err
}

func (s *PedidoService) GetAllDetallePedidosByIDPedido(ctx context.Context, idPedido int) ([]domain.DetallePedidoResponse, error) {
	var err error
	var detallePedido []domain.DetallePedidoResponse
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		detallePedido, err = s.repository.GetAllDetallePedidoByIDPedido(ctx, tx, idPedido)
		return err
	})
	return detallePedido, err
}

func (s *PedidoService) GetRankingDePedidosPorCliente(ctx context.Context, desde, hasta string) ([]domain.PedidosPorCliente, error) {
	var err error
	var pedidosByClient []domain.PedidosPorCliente
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		pedidosByClient, err = s.repository.GetRankingDePedidosPorCliente(ctx, tx, desde, hasta)
		return err
	})
	return pedidosByClient, err
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

func (s *PedidoService) UpdateEstadoPedido(ctx context.Context, estado, IDPedido int) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {

		pedido, err := s.repository.GetByID(ctx, tx, IDPedido)
		if err != nil {
			return err
		}

		fmt.Println("pedido", pedido)
		err = s.repository.UpdateEstadoPedido(ctx, tx, estado, IDPedido)
		if err != nil {
			return err
		}
		var descuento, costo_total float64
		//var hardcodeta string = "hardcodeta"
		if estado == domain.FACTURADO {
			// ok, err := s.repository.DescontarStock(ctx, tx, IDPedido, estado)
			// if !ok || err != nil {
			// 	return err
			// }
			costo_total, err = s.repository.GetCostoTotalByPedido(ctx, tx, IDPedido)
			if err != nil {
				return err
			}
			if pedido.TipoEnvio == domain.ENVIO_RETIRO_LOCAL {
				descuento = pedido.Total * 0.1
			}
			factura := domain.Factura{
				MontoDescuento: &descuento,
				FormaPago:      pedido.DetalleEnvio,
				TotalVenta:     &pedido.Total,
				TotalCosto:     &costo_total,
				IDPedido:       &IDPedido,
			}
			err = s.repositoryFactura.Insert(ctx, tx, factura)
			if err != nil {
				return err
			}
			fmt.Println("Factura:", factura)
		}
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

// func (s *PedidoService) AddPedido(ctx context.Context, pedido domain.Pedido) error {
// 	var err error
// 	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
// 		_, err = s.repository.Insert(ctx, tx, pedido)
// 		return err
// 	})
// 	return nil
// }

func (s *PedidoService) GenerarPedido(ctx context.Context, generarPedido domain.GenerarPedido) (domain.GenerarPedido, error) {
	var err error
	var pedido = generarPedido.Pedido
	var detallePedido = generarPedido.DetallePedido
	var idPedido, total, tiempoCocinaAcum, cantidadDeCocineros, tiempoTotalEstimado int
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {

		// controlamos el tiempo de demora de cada producto
		for _, det := range detallePedido {
			if det.TiempoEstimadoCocina != nil {
				tiempoCocinaAcum += *det.TiempoEstimadoCocina
				fmt.Println("tiempoCocinaAcum:", tiempoCocinaAcum)
			}
		}
		cantidadDeCocineros, err = s.repositoryLogin.CantidadDeCocineros(ctx, tx)
		if err != nil {
			return err
		}
		if pedido.TipoEnvio == domain.ENVIO_DELIVERY {
			tiempoTotalEstimado = (tiempoCocinaAcum / cantidadDeCocineros) + 10
		} else {
			tiempoTotalEstimado = (tiempoCocinaAcum / cantidadDeCocineros)
		}
		fmt.Println("cantidadDeCocineros service", cantidadDeCocineros)
		fmt.Println("tiempoCocinaAcum:", tiempoCocinaAcum)
		fmt.Println("tiempoTotalEstimado:", tiempoTotalEstimado)

		// creamos un 'pedido' en la BD y nos retorna el ID
		idPedido, err = s.repository.Insert(ctx, tx, pedido, tiempoTotalEstimado)
		fmt.Println("time.Now():", time.Now())

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
		if pedido.Estado != domain.PENDIENTE_APROBACION {
			return errors.New("solo se puede aceptar pedidos en estado 'pendiente de aprobacion' :: 1 ")
		}

		ok, err = s.repository.DescontarStock(ctx, tx, idPedido, pedido.Estado)
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

func (s *PedidoService) VerificarStock(ctx context.Context, idArticulo, amount int, esBebida bool) (bool, error) {
	var err error
	var ok bool = true
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		// obtengo el pedido por id para comprobar que exista
		if esBebida {
			fmt.Println("es bebida true", esBebida)
			ok, err = s.repository.VerificarStockBebidas(ctx, tx, idArticulo, amount)
		} else {
			fmt.Println("es bebida false", esBebida)
			ok, err = s.repository.VerificarStockManufacturado(ctx, tx, idArticulo, amount)
		}
		if !ok || err != nil {
			fmt.Println("err", err)
			return err
		}
		return err
	})
	fmt.Println(err)
	if err != nil {
		return false, err
	}
	return ok, err
}

func (s *PedidoService) RankingComidasMasPedidas(ctx context.Context, desde, hasta string) ([]domain.RankingComidasMasPedidas, error) {
	var err error
	var rankingComidasMasPedidas []domain.RankingComidasMasPedidas
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		rankingComidasMasPedidas, err = s.repository.RankingComidasMasPedidas(ctx, tx, desde, hasta)
		return err
	})
	return rankingComidasMasPedidas, err
}

func (s *PedidoService) CancelarPedido(ctx context.Context, idPedido int) error {
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.CancelarPedido(ctx, tx, idPedido)
		return err
	})
	return err
}
