package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IArticuloInsumoController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	GetAllCarritoCompleto(*gin.Context)
	AddArticuloInsumo(*gin.Context)
	UpdateArticuloInsumo(*gin.Context)
	DeleteArticuloInsumo(*gin.Context)
	AgregarStockInsumo(*gin.Context)
}

type ArticuloInsumoController struct {
	service services.IArticuloInsumoService
}

func NewArticuloInsumoController(service services.IArticuloInsumoService) *ArticuloInsumoController {
	return &ArticuloInsumoController{service}
}

func (c *ArticuloInsumoController) GetByID(ctx *gin.Context) {
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
	articuloInsumo, err := c.service.GetByID(ctx, ID)
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
	ctx.JSON(200, articuloInsumo)
}

func (c *ArticuloInsumoController) GetAll(ctx *gin.Context) {
	articulos, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, articulos)
}

func (c *ArticuloInsumoController) GetAllCarritoCompleto(ctx *gin.Context) {
	articulos, err := c.service.GetAllCarritoCompleto(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, articulos)
}

func (c *ArticuloInsumoController) AddArticuloInsumo(ctx *gin.Context) {
	var articuloInsumo domain.ArticuloInsumo

	err := ctx.BindJSON(&articuloInsumo)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}
	fmt.Println("articuloInsumo:", articuloInsumo)

	err = c.service.AddArticuloInsumo(ctx, articuloInsumo)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, gin.H{"articulo added successfully": articuloInsumo})
}

func (c *ArticuloInsumoController) UpdateArticuloInsumo(ctx *gin.Context) {
	var articulo domain.ArticuloInsumo

	err := ctx.BindJSON(&articulo)
	fmt.Println("articulo:", articulo)

	if err != nil {
		ctx.JSON(400, errors.New("invalid articulo to be updated"))
		return
	}

	err = c.service.UpdateArticuloInsumo(ctx, articulo)
	if err != nil {
		ctx.JSON(500, errors.New("internal error server"))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "articulo updated successfully",
	})
}

func (c *ArticuloInsumoController) DeleteArticuloInsumo(ctx *gin.Context) {
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

	err = c.service.DeleteArticuloInsumo(ctx, ID)
	if err != nil {
		if err.Error() == errNotFound.Error() {
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

	ctx.JSON(200, gin.H{
		"message": "factura deleted successfully",
	})
}

func (c *ArticuloInsumoController) AgregarStockInsumo(ctx *gin.Context) {
	var agregarStock domain.AgregarStockInsumo

	err := ctx.BindJSON(&agregarStock)
	fmt.Println("agregarStock:", agregarStock)

	if err != nil {
		ctx.JSON(400, errors.New("invalid articulo to be add stock"))
		return
	}

	err = c.service.SumarStockInsumo(ctx, agregarStock)
	if err != nil {
		ctx.JSON(500, errors.New("internal error server"))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "articulo updated successfully",
	})
}
