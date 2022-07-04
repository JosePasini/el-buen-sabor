package controllers

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IArticuloManufacturadoController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	AddArticuloManufacturado(*gin.Context)
	UpdateArticuloManufacturado(*gin.Context)
	DeleteArticuloManufacturado(*gin.Context)
}

type ArticuloManufacturadoController struct {
	service services.IArticuloManufacturadoService
}

func NewArticuloManufacturadoController(service services.IArticuloManufacturadoService) *ArticuloManufacturadoController {
	return &ArticuloManufacturadoController{service}
}

func (c *ArticuloManufacturadoController) GetByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloManufacturadoController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloManufacturadoController) AddArticuloManufacturado(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})

}
func (c *ArticuloManufacturadoController) UpdateArticuloManufacturado(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})

}
func (c *ArticuloManufacturadoController) DeleteArticuloManufacturado(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})

}
