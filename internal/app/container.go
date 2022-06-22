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

	// Repositorys
	InstrumentoRepository domain.IInstrumentoRepository
}

func NewContainer(config instrumentos.AppConfig, db database.DB) Container {
	instrumentoRepository := storage.NewMySQLInstrumentoRepository()
	instrumentoService := services.NewInstrumentoService(db, instrumentoRepository)

	return Container{
		Config:                config,
		InstrumentoService:    instrumentoService,
		InstrumentoRepository: instrumentoRepository,
	}
}
