package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

var (
	errNotFound = errors.New("factura not found")
	errInternal = errors.New("internal server error")
)

type IFacturaController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	GetAllByCliente(*gin.Context)
	GetByIDPedido(*gin.Context)
	AddFactura(*gin.Context)
	UpdateFactura(*gin.Context)
	DeleteFactura(*gin.Context)
	RecaudacionesDiarias(*gin.Context)
	RecaudacionesMensuales(*gin.Context)
	RecaudacionesPeriodoTiempo(*gin.Context)
	ObtenerGanancias(*gin.Context)
}

type FacturaController struct {
	service services.IFacturaService
}

func NewFacturaController(service services.IFacturaService) *FacturaController {
	return &FacturaController{service}
}

func (c *FacturaController) AddFactura(ctx *gin.Context) {
	var factura domain.Factura
	err := ctx.BindJSON(&factura)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}
	fmt.Println("AddFactura controller:", factura)

	err = c.service.AddFactura(ctx, factura)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, factura)
}

func (c *FacturaController) GetByIDPedido(ctx *gin.Context) {
	idPedido := ctx.Param("idPedido")

	ID, err := strconv.Atoi(idPedido)
	if err != nil {
		ctx.JSON(400, errors.New("invalid pedido id"))
		return
	}
	factura, err := c.service.GetByIDPedido(ctx, ID)
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
	ctx.JSON(200, factura)
}

func (c *FacturaController) GetByID(ctx *gin.Context) {
	idFactura := ctx.Param("idFactura")

	ID, err := strconv.Atoi(idFactura)
	if err != nil {
		ctx.JSON(400, errors.New("invalid factura id"))
		return
	}
	factura, err := c.service.GetByID(ctx, ID)
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
	ctx.JSON(200, factura)
}

func (c *FacturaController) GetAll(ctx *gin.Context) {
	facturas, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, facturas)
}
func (c *FacturaController) GetAllByCliente(ctx *gin.Context) {
	idParam := ctx.Param("idCliente")

	idCliente, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("invalid factura id"))
		return
	}
	facturas, err := c.service.GetAllByCliente(ctx, idCliente)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, facturas)
}

func (c *FacturaController) RecaudacionesDiarias(ctx *gin.Context) {
	fecha := ctx.Query("fecha")
	recaudaciones, err := c.service.RecaudacionesDiarias(ctx, fecha)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, recaudaciones)
}

func (c *FacturaController) RecaudacionesMensuales(ctx *gin.Context) {
	month := ctx.Query("month")
	year := ctx.Query("year")
	recaudaciones, err := c.service.RecaudacionesMensuales(ctx, month, year)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, recaudaciones)
}

func (c *FacturaController) RecaudacionesPeriodoTiempo(ctx *gin.Context) {
	desde := ctx.Query("desde")
	hasta := ctx.Query("hasta")
	recaudaciones, err := c.service.RecaudacionesPeriodoTiempo(ctx, desde, hasta)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, recaudaciones)
}

func (c *FacturaController) ObtenerGanancias(ctx *gin.Context) {
	desde := ctx.Query("desde")
	hasta := ctx.Query("hasta")
	ganancias, err := c.service.ObtenerGanancias(ctx, desde, hasta)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, ganancias)
}

func (c *FacturaController) DeleteFactura(ctx *gin.Context) {
	idFactura := ctx.Param("idFactura")
	if idFactura == "" {
		ctx.JSON(400, errors.New("invalid factura"))
		return
	}

	ID, err := strconv.Atoi(idFactura)
	if err != nil {
		ctx.JSON(400, errors.New("id factura must be a number"))
		return
	}

	err = c.service.DeleteFactura(ctx, ID)
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

func (c *FacturaController) UpdateFactura(ctx *gin.Context) {
	var factura domain.Factura

	err := ctx.BindJSON(&factura)
	fmt.Println("Factura:", factura)

	if err != nil {
		ctx.JSON(400, errors.New("invalid factura to be updated"))
		return
	}

	err = c.service.UpdateFactura(ctx, factura)
	if err != nil {
		ctx.JSON(500, errors.New("internal error server"))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "factura updated successfully",
	})
}
