package services

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/storage"
)

type IArticuloManufacturadoDetalleService interface {
	// GetAll(context.Context) ([]domain.Factura, error)
	// GetByID(context.Context, int) (*domain.Factura, error)
	// UpdateFactura(context.Context, domain.Factura) error
	// DeleteFactura(context.Context, int) error
	// AddFactura(context.Context, domain.Factura) error
}

type ArticuloManufacturadoDetalleService struct {
	db         database.DB
	repository storage.IArticuloManufacturadoDetalleRepository
}

func NewArticuloManufacturadoDetalleService(db database.DB, repository storage.IArticuloManufacturadoDetalleRepository) *ArticuloManufacturadoDetalleService {
	return &ArticuloManufacturadoDetalleService{db, repository}
}
