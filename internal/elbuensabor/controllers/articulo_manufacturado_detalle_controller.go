package controllers

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IArticuloManufacturadoDetalleController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	AddArticuloManufacturadoDetalle(*gin.Context)
	UpdateArticuloManufacturadoDetalle(*gin.Context)
	DeleteArticuloManufacturadoDetalle(*gin.Context)
}

type ArticuloManufacturadoDetalleController struct {
	service services.IArticuloManufacturadoDetalleService
}

func NewArticuloManufacturadoDetalleController(service services.IArticuloManufacturadoDetalleService) *ArticuloManufacturadoDetalleController {
	return &ArticuloManufacturadoDetalleController{service}
}

func (c *ArticuloManufacturadoDetalleController) GetByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloManufacturadoDetalleController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloManufacturadoDetalleController) AddArticuloManufacturadoDetalle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})

}
func (c *ArticuloManufacturadoDetalleController) UpdateArticuloManufacturadoDetalle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloManufacturadoDetalleController) DeleteArticuloManufacturadoDetalle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})

}
