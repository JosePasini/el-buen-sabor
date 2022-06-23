package controllers

import (
	"errors"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	AddUsuario(ctx *gin.Context)
	LoginUsuario(ctx *gin.Context)
}

type LoginController struct {
	service services.ILoginService
}

func NewLoginController(service services.ILoginService) *LoginController {
	return &LoginController{service}
}

func (c *LoginController) AddUsuario(ctx *gin.Context) {
	var usuario domain.Usuario
	err := ctx.BindJSON(&usuario)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}
	fmt.Println("Usuario:", *usuario.Hash, *usuario.Nombre, *usuario.Apellido, *usuario.Usuario, *usuario.Mail)
	err = c.service.AddUsuario(ctx, usuario)
	if err != nil {
		ctx.JSON(400, errors.New("error internal server"))
		return
	}
	ctx.JSON(200, gin.H{"message": "usuario agregado correctamente"})
}

func (c *LoginController) LoginUsuario(ctx *gin.Context) {
	var login domain.Login
	err := ctx.BindJSON(&login)
	if err != nil {
		ctx.JSON(400, errors.New("Error"))
		return
	}
	ok, err := c.service.LoginUsuario(ctx, login)
	if !ok || err != nil {
		fmt.Println("Login Incorrecto")
		fmt.Println("Login:", login)
		fmt.Println("Ok:", ok)
		ctx.JSON(400, errors.New("usuario o contraseña incorrectos"))
		return
	}
	fmt.Println("Login Correcto")
	fmt.Println("Login", login)
	fmt.Println("Ok:", ok)
	ctx.JSON(200, gin.H{"message": "crecendiales correctas"})
}
