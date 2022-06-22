package app

import (
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/database"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/domain"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/services"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/storage"
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
