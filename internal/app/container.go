package app

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
)

type Container struct {
	Config elbuensabor.AppConfig

	// Services
	FacturaService                      services.IFacturaService
	LoginService                        services.ILoginService
	PedidoService                       services.IPedidoService
	ArticuloManufacturadoDetalleService services.IArticuloManufacturadoDetalleService
	ArticuloManufacturadoService        services.IArticuloManufacturadoService
	ArticuloInsumoService               services.IArticuloInsumoService

	// Repositorys
	InstrumentoRepository                  domain.IFacturaRepository
	LoginRepository                        domain.ILoginRepository
	PedidoRepository                       domain.IPedidoRepository
	ArticuloManufacturadoDetalleRepository storage.IArticuloManufacturadoDetalleRepository
	ArticuloManufacturadoRepository        storage.IArticuloManufacturadoRepository
	ArticuloInsumoRepository               storage.IArticuloInsumoRepository
}

func NewContainer(config elbuensabor.AppConfig, db database.DB) Container {
	facturaRepository := storage.NewMySQLInstrumentoRepository()
	facturaService := services.NewFacturaService(db, facturaRepository)

	loginRepository := storage.NewMySQLLoginRepository()
	loginService := services.NewLoginService(db, loginRepository)

	pedidoRepository := storage.NewMySQLPedidoRepository()
	pedidoService := services.NewPedidoService(db, pedidoRepository)

	articuloManufacturadoDetalleRepository := storage.NewMySQLArticuloManufacturadoDetalleRepository()
	articuloManufacturadoDetalleService := services.NewArticuloManufacturadoDetalleService(db, articuloManufacturadoDetalleRepository)

	articuloManufacturadoRepository := storage.NewMySQLArticuloManufacturadoRepository()
	articuloManufacturadoService := services.NewArticuloManufacturadoService(db, articuloManufacturadoRepository)

	articuloInsumoRepository := storage.NewMySQLArticuloInsumoRepository()
	articuloInsumoService := services.NewArticuloInsumoService(db, articuloInsumoRepository)

	return Container{
		Config:                config,
		FacturaService:        facturaService,
		InstrumentoRepository: facturaRepository,

		LoginService:    loginService,
		LoginRepository: loginRepository,

		PedidoService:    pedidoService,
		PedidoRepository: pedidoRepository,

		ArticuloManufacturadoDetalleService:    articuloManufacturadoDetalleService,
		ArticuloManufacturadoDetalleRepository: articuloManufacturadoDetalleRepository,

		ArticuloManufacturadoService:    articuloManufacturadoService,
		ArticuloManufacturadoRepository: articuloManufacturadoRepository,

		ArticuloInsumoService:    articuloInsumoService,
		ArticuloInsumoRepository: articuloInsumoRepository,
	}
}
