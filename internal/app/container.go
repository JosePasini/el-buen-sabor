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
	InstrumentoService services.IInstrumentoService
	LoginService       services.ILoginService

	// Repositorys
	InstrumentoRepository domain.IInstrumentoRepository
	LoginRepository       domain.ILoginRepository
}

func NewContainer(config elbuensabor.AppConfig, db database.DB) Container {
	instrumentoRepository := storage.NewMySQLInstrumentoRepository()
	instrumentoService := services.NewInstrumentoService(db, instrumentoRepository)

	loginRepository := storage.NewMySQLLoginRepository()
	loginService := services.NewLoginService(db, loginRepository)

	return Container{
		Config:                config,
		InstrumentoService:    instrumentoService,
		InstrumentoRepository: instrumentoRepository,

		LoginService:    loginService,
		LoginRepository: loginRepository,
	}
}
