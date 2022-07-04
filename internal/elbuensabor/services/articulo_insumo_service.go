package services

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
)

type IArticuloInsumoService interface {
	// GetAll(context.Context) ([]domain.Factura, error)
	// GetByID(context.Context, int) (*domain.Factura, error)
	// UpdateFactura(context.Context, domain.Factura) error
	// DeleteFactura(context.Context, int) error
	// AddFactura(context.Context, domain.Factura) error
}

type ArticuloInsumoService struct {
	db         database.DB
	repository storage.IArticuloInsumoRepository
}

func NewArticuloInsumoService(db database.DB, repository storage.IArticuloInsumoRepository) *ArticuloInsumoService {
	return &ArticuloInsumoService{db, repository}
}
