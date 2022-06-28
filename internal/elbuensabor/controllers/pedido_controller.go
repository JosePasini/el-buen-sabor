package controllers

import (
	"errors"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IPedidoController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	AddPedido(*gin.Context)
	UpdatePedido(*gin.Context)
	DeletePedido(*gin.Context)
}

type PedidoController struct {
	service services.IPedidoService
}

func NewPedidoController(service services.IPedidoService) *PedidoController {
	return &PedidoController{service}
}

func (c PedidoController) GetByID(ctx *gin.Context) {
	pedidoID := ctx.Param("idPedido")

	if pedidoID == "" {
		ctx.JSON(400, errors.New("invalid pedido"))
		return
	}

	ID, err := strconv.Atoi(pedidoID)
	if err != nil {
		ctx.JSON(400, errors.New("invalid pedido id"))
		return
	}
	pedido, err := c.service.GetByID(ctx, ID)
	if err != nil {
		if err.Error() == errInternal.Error() {
			ctx.JSON(404, gin.H{
				"message": "pedido not found",
			})
			return
		}
	}
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, pedido)
}

func (c PedidoController) GetAll(ctx *gin.Context) {
	pedidos, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}
	ctx.JSON(200, pedidos)
}

func (c PedidoController) AddPedido(ctx *gin.Context) {
	var pedido domain.Pedido
	err := ctx.BindJSON(&pedido)
	if err != nil {
		ctx.JSON(400, errors.New("error bind"))
		return
	}

	err = c.service.AddPedido(ctx, pedido)
	if err != nil {
		ctx.JSON(400, errors.New("add pedido error"))
		return
	}
	ctx.JSON(200, gin.H{"status": 200, "pedido": pedido})
}
func (c PedidoController) UpdatePedido(ctx *gin.Context) {
	var pedido domain.Pedido

	err := ctx.BindJSON(&pedido)

	if err != nil {
		ctx.JSON(400, errors.New("invalid pedido to be updated"))
		return
	}

	err = c.service.UpdatePedido(ctx, pedido)
	if err != nil {
		ctx.JSON(500, errors.New("internal error server"))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "pedido updated successfully",
	})
}

func (c PedidoController) DeletePedido(ctx *gin.Context) {
	idPedido := ctx.Param("idPedido")
	if idPedido == "" {
		ctx.JSON(400, errors.New("invalid instrument"))
		return
	}

	ID, err := strconv.Atoi(idPedido)
	if err != nil {
		ctx.JSON(400, errors.New("id instrument must be a number"))
		return
	}

	err = c.service.DeletePedido(ctx, ID)
	if err != nil {
		if err.Error() == errNotFound.Error() {
			ctx.JSON(404, gin.H{
				"message": "pedido not found",
			})
			return
		}
	}

	if err != nil {
		ctx.JSON(500, errors.New("Error internal server error: "+err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "pedido deleted successfully",
	})
}
