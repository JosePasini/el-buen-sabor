package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IArticuloManufacturadoController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	GetAllAvailable(*gin.Context)
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
	idParam := ctx.Param("id")

	if idParam == "" {
		ctx.JSON(400, errors.New("invalid id param"))
		return
	}

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("invalid factura id"))
		return
	}
	articuloManufacturado, err := c.service.GetByID(ctx, ID)
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
	ctx.JSON(200, articuloManufacturado)
}

func (c *ArticuloManufacturadoController) GetAll(ctx *gin.Context) {
	articulosManufacturados, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, articulosManufacturados)
}

func (c *ArticuloManufacturadoController) GetAllAvailable(ctx *gin.Context) {
	articulosManufacturados, err := c.service.GetAllAvailable(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, articulosManufacturados)
}

func (c *ArticuloManufacturadoController) AddArticuloManufacturado(ctx *gin.Context) {
	var articuloManufacturado domain.ArticuloManufacturado

	err := ctx.BindJSON(&articuloManufacturado)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}
	fmt.Println("articuloManufacturado:", articuloManufacturado)

	err = c.service.AddArticuloManufacturado(ctx, articuloManufacturado)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, gin.H{"articulo added successfully": articuloManufacturado})
}

func (c *ArticuloManufacturadoController) UpdateArticuloManufacturado(ctx *gin.Context) {
	var articuloManufacturado domain.ArticuloManufacturado

	err := ctx.BindJSON(&articuloManufacturado)
	fmt.Println("articulo:", articuloManufacturado)

	if err != nil {
		ctx.JSON(400, errors.New("invalid articuloManufacturado to be updated"))
		return
	}

	err = c.service.UpdateArticuloManufacturado(ctx, articuloManufacturado)
	if err != nil {
		ctx.JSON(500, errors.New("internal error server"))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "articuloManufacturado updated successfully",
	})
}

func (c *ArticuloManufacturadoController) DeleteArticuloManufacturado(ctx *gin.Context) {

	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(400, errors.New("invalid factura"))
		return
	}

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("id factura must be a number"))
		return
	}

	err = c.service.DeleteArticuloManufacturado(ctx, ID)
	if err != nil {
		if err.Error() == errNotFound.Error() {
			ctx.JSON(404, gin.H{
				"message": "Articulo manufacturado not found",
			})
			return
		}
	}

	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Articulo manufacturado deleted successfully",
	})
}
