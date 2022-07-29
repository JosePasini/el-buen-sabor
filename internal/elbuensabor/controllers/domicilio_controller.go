package controllers

import (
	"errors"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IDomicilioController interface {
	AddDomicilio(*gin.Context)
	UpdateDomicilio(*gin.Context)
	GetAllDomicilioByUsuario(*gin.Context)
}

type DomicilioController struct {
	service services.IDomicilioService
}

func NewDomicilioController(service services.IDomicilioService) *DomicilioController {
	return &DomicilioController{service}
}

func (c *DomicilioController) AddDomicilio(ctx *gin.Context) {
	var domicilio domain.Domicilio
	err := ctx.BindJSON(&domicilio)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}

	err = c.service.AddDomicilio(ctx, domicilio)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, domicilio)
}

func (c *DomicilioController) UpdateDomicilio(ctx *gin.Context) {
	var domicilio domain.Domicilio
	err := ctx.BindJSON(&domicilio)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}

	err = c.service.UpdateDomicilio(ctx, domicilio)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, domicilio)
}

func (c *DomicilioController) GetAllDomicilioByUsuario(ctx *gin.Context) {

	idParam := ctx.Param("idUsuario")

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("invalid user id"))
		return
	}
	domiciliosByUsuarios, err := c.service.GetAllDomicilioByUsuario(ctx, ID)
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
	ctx.JSON(200, domiciliosByUsuarios)
}
