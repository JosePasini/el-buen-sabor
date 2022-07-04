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
	AddFactura(*gin.Context)
	UpdateFactura(*gin.Context)
	DeleteFactura(*gin.Context)
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

func (c *FacturaController) GetByID(ctx *gin.Context) {
	idFactura := ctx.Param("idFactura")

	if idFactura == "" {
		ctx.JSON(400, errors.New("invalid factura"))
		return
	}

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
