package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
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
	idParam := ctx.Param("id")

	if idParam == "" {
		ctx.JSON(400, errors.New("invalid id param"))
		return
	}

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("invalid articulo manufacturado detalle id"))
		return
	}
	articuloManufacturadoDetalle, err := c.service.GetByID(ctx, ID)
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
	ctx.JSON(200, articuloManufacturadoDetalle)
}

func (c *ArticuloManufacturadoDetalleController) GetAll(ctx *gin.Context) {
	articulosManufacturadosDetalles, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, articulosManufacturadosDetalles)
}

func (c *ArticuloManufacturadoDetalleController) AddArticuloManufacturadoDetalle(ctx *gin.Context) {
	var articuloManufacturadoDetalle domain.ArticuloManufacturadoDetalle

	err := ctx.BindJSON(&articuloManufacturadoDetalle)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}
	fmt.Println("articuloManufacturadoDetalle:", articuloManufacturadoDetalle)

	err = c.service.AddArticuloManufacturadoDetalle(ctx, articuloManufacturadoDetalle)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, gin.H{"articulo manufacturado detalle added successfully": articuloManufacturadoDetalle})
}

func (c *ArticuloManufacturadoDetalleController) UpdateArticuloManufacturadoDetalle(ctx *gin.Context) {
	var articuloManufacturado domain.ArticuloManufacturadoDetalle

	err := ctx.BindJSON(&articuloManufacturado)
	fmt.Println("articuloManufacturado:", articuloManufacturado)

	if err != nil {
		ctx.JSON(400, errors.New("invalid articulo manufacturado detalle to be updated"))
		return
	}

	err = c.service.UpdateArticuloManufacturadoDetalle(ctx, articuloManufacturado)
	if err != nil {
		ctx.JSON(500, errors.New("internal error server"))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "articuloManufacturado detalle updated successfully",
	})
}

func (c *ArticuloManufacturadoDetalleController) DeleteArticuloManufacturadoDetalle(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(400, errors.New("invalid articulo manufacturado detalle"))
		return
	}

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("id articulo manufacturado detalle must be a number"))
		return
	}

	err = c.service.DeleteArticuloManufacturadoDetalle(ctx, ID)
	if err != nil {
		if err.Error() == errNotFound.Error() {
			ctx.JSON(404, gin.H{
				"message": "Articulo manufacturado detalle not found",
			})
			return
		}
	}

	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Articulo manufacturado detalle deleted successfully",
	})
}
