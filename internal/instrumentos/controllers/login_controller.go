package controllers

import (
	"errors"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/services"
	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	AddUsuario(ctx *gin.Context)
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
	fmt.Println("Usuario:", usuario.Nombre, usuario.Apellido)
	fmt.Println("Usuario:", *usuario.Nombre, *usuario.Apellido)
	// err = c.service.AddUsuario(ctx, usuario)
	// if err != nil {
	// 	ctx.JSON(400, errors.New("Error Internal Server"))
	// 	return
	// }
	// ctx.JSON(200, nil)
}
