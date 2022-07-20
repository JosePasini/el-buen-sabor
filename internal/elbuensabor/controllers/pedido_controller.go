package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type IPedidoController interface {
	GetByID(*gin.Context)
	GetAll(*gin.Context)
	AddPedido(*gin.Context)
	AceptarPedido(*gin.Context)
	GenerarPedido(*gin.Context)
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

func (c PedidoController) GenerarPedido(ctx *gin.Context) {
	var pedido domain.GenerarPedido

	fmt.Println("1")

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &pedido)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	fmt.Println("2")
	fmt.Println(pedido)
	pedido, err = c.service.GenerarPedido(ctx, pedido)
	if err != nil {
		ctx.JSON(400, errors.New("generate pedido error"))
		return
	}
	fmt.Println("3")

	ctx.JSON(200, gin.H{"status": 200, "pedido": pedido})
}

func (c PedidoController) AceptarPedido(ctx *gin.Context) {
	fmt.Println(" -- principio --")
	idParam := ctx.Param("idPedido")

	fmt.Println(idParam)
	idPedido, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("add pedido error"))
		return
	}
	ok, err := c.service.AceptarPedido(ctx, idPedido)
	if err != nil || !ok {
		ctx.JSON(400, errors.New("aceptar pedido error"))
		return
	}
	fmt.Println(" -- fin -- ")
	ctx.JSON(200, gin.H{"status": 200, "ok": ok})
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
