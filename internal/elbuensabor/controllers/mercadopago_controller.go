package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/eduardo-mior/mercadopago-sdk-go"
	"github.com/gin-gonic/gin"
)

type IMercadoPagoController interface {
	Pagar(*gin.Context)
	MetodosDePago(*gin.Context)
}

type MercadoPagoController struct {
}

func NewMercadoPagoController() *MercadoPagoController {
	return &MercadoPagoController{}
}

func (c *MercadoPagoController) Pagar(ctx *gin.Context) {
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
		Email: cliente.Email,
		Name:  cliente.Nombre,
	}
	fmt.Println("clienteMP", clienteMP)
	BACK_URL_MP := mercadopago.BackUrls{
		Success: "https://frontprueba.herokuapp.com/pedir",
		Pending: "https://frontprueba.herokuapp.com/pedir",
		Failure: "https://frontprueba.herokuapp.com/pedir",
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

func (c *MercadoPagoController) MetodosDePago(ctx *gin.Context) {
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
