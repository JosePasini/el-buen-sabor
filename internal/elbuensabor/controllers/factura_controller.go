package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/eduardo-mior/mercadopago-sdk-go"
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
	MercadoPago(*gin.Context)
	MetodosDePago(*gin.Context)
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

func (c *FacturaController) MercadoPago(ctx *gin.Context) {
	TOKEN_EL_BUEN_SABOR_TEST := os.Getenv("TOKEN_EL_BUEN_SABOR_TEST")

	type RequestMercadoPago struct {
		ProductsList []domain.ProductoMercadoPago `json:"producto_mercado_pago"`
		Cliente      domain.UsuarioMercadoPago    `json:"usuario"`
	}
	var requestMercadoPago RequestMercadoPago

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &requestMercadoPago)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	listaProductos := requestMercadoPago.ProductsList
	cliente := requestMercadoPago.Cliente

	fmt.Println("listaProductos", listaProductos)
	fmt.Println("cliente", cliente)

	var itemsMercadoPago []mercadopago.Item
	for _, prod := range listaProductos {
		itemMP := mercadopago.Item{
			Title:      prod.Denominacion,
			PictureURL: prod.Imagen,
			Quantity:   float64(prod.Amount),
			UnitPrice:  prod.PrecioVenta,
		}
		itemsMercadoPago = append(itemsMercadoPago, itemMP)
	}
	clienteMP := mercadopago.Payer{
		// El email tiene que ser de test y matchear con una cuenta de MP TEST, sino no funciona.
		// Email: "test_user_53114826@testuser.com",
		Email: cliente.Email,
		Name:  cliente.Nombre,
		// Identification: mercadopago.PayerIdentification{
		// 	Type:   cliente.TypeIdentification,
		// 	Number: cliente.NumberIdentification,
		// },
	}
	fmt.Println("clienteMP", clienteMP)
	BACK_URL_MP := mercadopago.BackUrls{
		Success: "localhost:3000/pedir",
		Pending: "localhost:3000/pedir",
		Failure: "localhost:3000/pedir",
	}

	paymentResponse, mercadopagoErr, err := mercadopago.CreatePayment(mercadopago.PaymentRequest{
		Items:      itemsMercadoPago,
		Payer:      clienteMP,
		BackUrls:   BACK_URL_MP,
		AutoReturn: "approved",
	}, TOKEN_EL_BUEN_SABOR_TEST)

	if err != nil {
		ctx.JSON(400, errors.New("/mercado-pago error"))
		return
	} else if mercadopagoErr != nil {
		// Erro retornado do MercadoPago
		ctx.JSON(400, mercadopagoErr)
		return
	} else {
		// Sucesso!
		ctx.JSON(200, paymentResponse)
		return
	}
}

func (c *FacturaController) MetodosDePago(ctx *gin.Context) {
	TOKEN_EL_BUEN_SABOR_TEST := os.Getenv("TOKEN_EL_BUEN_SABOR_TEST")

	//identificationTypes, mercadopagoErr, err := mercadopago.GetIdentificationTypes(TOKEN_EL_BUEN_SABOR_TEST)
	metodosPago, mercadopagoErr, err := mercadopago.GetPaymentMethods(TOKEN_EL_BUEN_SABOR_TEST)

	if err != nil {
		// Erro inesperado
		ctx.JSON(400, err)
	} else if mercadopagoErr != nil {
		// Erro retornado do MercadoPago
		ctx.JSON(400, mercadopagoErr)
	} else {
		// Sucesso!
		ctx.JSON(200, metodosPago)
	}
}
