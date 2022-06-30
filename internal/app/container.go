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
	FacturaService services.IFacturaService
	LoginService   services.ILoginService
	PedidoService  services.IPedidoService

	// Repositorys
	InstrumentoRepository domain.IFacturaRepository
	LoginRepository       domain.ILoginRepository
	PedidoRepository      domain.IPedidoRepository
}

func NewContainer(config elbuensabor.AppConfig, db database.DB) Container {
	facturaRepository := storage.NewMySQLInstrumentoRepository()
	facturaService := services.NewFacturaService(db, facturaRepository)

	loginRepository := storage.NewMySQLLoginRepository()
	loginService := services.NewLoginService(db, loginRepository)

	pedidoRepository := storage.NewMySQLPedidoRepository()
	pedidoService := services.NewPedidoService(db, pedidoRepository)
	//pedidoService := services.NewPedidoService(db, pedidoRepository)

	return Container{
		Config:                config,
		FacturaService:        facturaService,
		InstrumentoRepository: facturaRepository,

		LoginService:    loginService,
		LoginRepository: loginRepository,

		PedidoService:    pedidoService,
		PedidoRepository: pedidoRepository,
	}
}
