package app

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/services"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/storage"
)

type Container struct {
	Config instrumentos.AppConfig

	// Services
	InstrumentoService services.IInstrumentoService
	LoginService       services.ILoginService

	// Repositorys
	InstrumentoRepository domain.IInstrumentoRepository
	LoginRepository       domain.ILoginRepository
}

func NewContainer(config instrumentos.AppConfig, db database.DB) Container {
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
