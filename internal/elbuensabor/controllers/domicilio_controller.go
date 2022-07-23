package controllers

import (
	"errors"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IDomicilioController interface {
	AddDomicilio(*gin.Context)
	UpdateDomicilio(*gin.Context)
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
