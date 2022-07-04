package controllers

import (
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IArticuloInsumoController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	AddArticuloInsumo(*gin.Context)
	UpdateArticuloInsumo(*gin.Context)
	DeleteArticuloInsumo(*gin.Context)
}

type ArticuloInsumoController struct {
	service services.IArticuloInsumoService
}

func NewArticuloInsumoController(service services.IArticuloInsumoService) *ArticuloInsumoController {
	return &ArticuloInsumoController{service}
}

func (c *ArticuloInsumoController) GetByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloInsumoController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloInsumoController) AddArticuloInsumo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloInsumoController) UpdateArticuloInsumo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
func (c *ArticuloInsumoController) DeleteArticuloInsumo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
