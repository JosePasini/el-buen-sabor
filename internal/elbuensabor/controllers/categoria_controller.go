package controllers

import (
	"errors"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type ICategoriaController interface {
	AddCategoria(*gin.Context)
	GetAllCategoria(*gin.Context)
}

type CategoriaController struct {
	service services.ICategoriaService
}

func NewCategoriaController(service services.ICategoriaService) *CategoriaController {
	return &CategoriaController{service}
}

func (c *CategoriaController) AddCategoria(ctx *gin.Context) {
	var categoria domain.Categoria
	err := ctx.BindJSON(&categoria)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}

	err = c.service.AddCategoria(ctx, categoria)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, categoria)
}

func (c *CategoriaController) GetAllCategoria(ctx *gin.Context) {
	categorias, err := c.service.GetAll(ctx)
	if err != nil {
		if err.Error() == errInternal.Error() {
			ctx.JSON(404, gin.H{
				"message": "factura not found",
			})
			return
		}
	}
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, categorias)
}
