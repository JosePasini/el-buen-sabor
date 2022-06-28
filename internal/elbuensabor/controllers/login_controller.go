package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	AddUsuario(ctx *gin.Context)
	LoginUsuario(ctx *gin.Context)
	GetAllUsuarios(ctx *gin.Context)
	GetUsuarioByID(ctx *gin.Context)
	DeleteUsuarioByID(ctx *gin.Context)
	UpdateUsuario(ctx *gin.Context)
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
	fmt.Println("Usuario:", *usuario.Hash, *usuario.Nombre, *usuario.Apellido, *usuario.Usuario, *usuario.Email)
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
		ctx.JSON(401, gin.H{"message": "crecendiales incorrectas"})
		return
	}
	fmt.Println("Login Correcto")
	fmt.Println("Login", login)
	fmt.Println("Ok:", ok)
	ctx.JSON(200, gin.H{"message": "crecendiales correctas"})
}

func (c *LoginController) GetAllUsuarios(ctx *gin.Context) {
	var listaUsuarios []domain.UsuarioResponse
	var err error

	listaUsuarios, err = c.service.GetAllUsuarios(ctx)
	if err != nil {
		fmt.Println("Login Incorrecto")
		ctx.JSON(400, errors.New("usuario o contraseña incorrectos"))
		return
	}
	ctx.JSON(200, listaUsuarios)
}

func (c *LoginController) GetUsuarioByID(ctx *gin.Context) {
	var usuario domain.UsuarioResponse
	var err error

	idParam := ctx.Param("id")

	if idParam == "" {
		ctx.JSON(400, errors.New("error get usuario by id"))
		return
	}

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("error get usuario by id"))
		return
	}

	usuario, err = c.service.GetUsuarioByID(ctx, ID)
	if err != nil {
		ctx.JSON(400, errors.New("error get usuario by id"))
		return
	}
	ctx.JSON(200, usuario)
}

func (c *LoginController) DeleteUsuarioByID(ctx *gin.Context) {
	var err error
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(400, errors.New("invalid id usuario"))
		return
	}

	ID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, errors.New("invalid usuario id"))
		return
	}

	ok, err := c.service.DeleteUsuarioByID(ctx, ID)
	if err != nil {
		fmt.Println("Error Delete Usuario")
		ctx.JSON(400, errors.New("error delete usuario"))
		return
	}

	if ok {
		ctx.JSON(200, ok)
		return
	}
	ctx.JSON(204, ok)
}

func (c *LoginController) UpdateUsuario(ctx *gin.Context) {
	var usuario domain.Usuario
	var err error

	err = ctx.BindJSON(&usuario)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "error update usuario by id, check the JSON"})
		return
	}

	usuario, err = c.service.UpdateUsuario(ctx, usuario)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "error in update usuario"})
		return
	}
	ctx.JSON(200, usuario)
}
