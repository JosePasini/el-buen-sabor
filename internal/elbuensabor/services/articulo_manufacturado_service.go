package services

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
)

type IArticuloManufacturadoService interface {
	// GetAll(context.Context) ([]domain.Factura, error)
	// GetByID(context.Context, int) (*domain.Factura, error)
	// UpdateFactura(context.Context, domain.Factura) error
	// DeleteFactura(context.Context, int) error
	// AddFactura(context.Context, domain.Factura) error
}

type ArticuloManufacturadoService struct {
	db         database.DB
	repository storage.IArticuloManufacturadoRepository
}

func NewArticuloManufacturadoService(db database.DB, repository storage.IArticuloManufacturadoRepository) *ArticuloManufacturadoService {
	return &ArticuloManufacturadoService{db, repository}
}
