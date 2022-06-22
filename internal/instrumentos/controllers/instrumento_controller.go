package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/domain"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/services"
	"github.com/gin-gonic/gin"
)

var (
	errNotFound = errors.New("instrument not found")
	errInternal = errors.New("internal server error")
)

type IInstrumentoController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	AddInstrument(*gin.Context)
	UpdateInstrument(*gin.Context)
	DeleteInstrument(*gin.Context)
}

type InstrumentoController struct {
	service services.IInstrumentoService
}

func NewInstrumentoController(service services.IInstrumentoService) *InstrumentoController {
	return &InstrumentoController{service}
}

func (c *InstrumentoController) AddInstrument(ctx *gin.Context) {
	var instrumento domain.Instrumento
	err := ctx.BindJSON(&instrumento)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}

	err = c.service.AddInstrumento(ctx, instrumento)
	if err != nil {
		ctx.JSON(400, errors.New("Error Internal Server"))
		return
	}
	ctx.JSON(200, instrumento)
}

func (c *InstrumentoController) GetByID(ctx *gin.Context) {
	instrumentoId := ctx.Param("idInstrumento")

	if instrumentoId == "" {
		ctx.JSON(400, errors.New("invalid instrument"))
		return
	}

	ID, err := strconv.Atoi(instrumentoId)
	if err != nil {
		ctx.JSON(400, errors.New("invalid instrument id"))
		return
	}
	instrumento, err := c.service.GetByID(ctx, ID)
	if err != nil {
		if err.Error() == errInternal.Error() {
			ctx.JSON(404, gin.H{
				"message": "instrument not found",
			})
			return
		}
	}
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, instrumento)

}

func (c *InstrumentoController) GetAll(ctx *gin.Context) {
	instruments, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, instruments)
}

func (c *InstrumentoController) DeleteInstrument(ctx *gin.Context) {
	idInstrumento := ctx.Param("idInstrumento")
	if idInstrumento == "" {
		ctx.JSON(400, errors.New("invalid instrument"))
		return
	}

	ID, err := strconv.Atoi(idInstrumento)
	if err != nil {
		ctx.JSON(400, errors.New("id instrument must be a number"))
		return
	}

	err = c.service.DeleteInstrument(ctx, ID)
	if err != nil {
		if err.Error() == errNotFound.Error() {
			ctx.JSON(404, gin.H{
				"message": "instrument not found",
			})
			return
		}
	}

	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "instrument deleted successfully",
	})
}

func (c *InstrumentoController) UpdateInstrument(ctx *gin.Context) {
	var instrument domain.Instrumento

	err := ctx.BindJSON(&instrument)
	fmt.Println("Instrumento:", instrument)
	fmt.Println("Instrumento:", *instrument.Instrumento)
	fmt.Println("Marca:", *instrument.Marca)

	if err != nil {
		ctx.JSON(400, errors.New("invalid instrument to be updated"))
		return
	}

	err = c.service.UpdateInstrument(ctx, instrument)
	if err != nil {
		ctx.JSON(500, errors.New("internal error server"))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "instrument updated successfully",
	})
}
