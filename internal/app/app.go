package app

import (
	"net/http"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/controllers"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	db     database.DB
	Config elbuensabor.AppConfig

	LoginService    services.ILoginService
	LoginController controllers.ILoginController

	PedidoService    services.IPedidoService
	PedidoController controllers.IPedidoController

	FacturaService    services.IFacturaService
	FacturaController controllers.IFacturaController

	ArticuloManufacturadoDetalleService    services.IArticuloManufacturadoDetalleService
	ArticuloManufacturadoDetalleController controllers.IArticuloManufacturadoDetalleController

	ArticuloManufacturadoService    services.IArticuloManufacturadoService
	ArticuloManufacturadoController controllers.IArticuloManufacturadoController

	ArticuloInsumoService    services.IArticuloInsumoService
	ArticuloInsumoController controllers.IArticuloInsumoController
}

func NewApp() (*App, error) {
	//scope := ("prod")
	scope := ("dev")

	config, err := NewConfig(scope)
	if err != nil {
		return &App{}, err
	}

	mysqlDB, err := database.NewMySQL(config.DB)
	if err != nil {
		return &App{}, err
	}

	container := NewContainer(config, mysqlDB)

	app := App{
		db:                                     mysqlDB,
		Config:                                 config,
		LoginService:                           container.LoginService,
		LoginController:                        controllers.NewLoginController(container.LoginService),
		PedidoService:                          container.PedidoService,
		PedidoController:                       controllers.NewPedidoController(container.PedidoService),
		FacturaService:                         container.FacturaService,
		FacturaController:                      controllers.NewFacturaController(container.FacturaService),
		ArticuloManufacturadoDetalleService:    container.ArticuloManufacturadoDetalleService,
		ArticuloManufacturadoDetalleController: controllers.NewArticuloManufacturadoDetalleController(container.ArticuloManufacturadoDetalleService),

		ArticuloManufacturadoService:    container.ArticuloManufacturadoService,
		ArticuloManufacturadoController: controllers.NewArticuloManufacturadoController(container.ArticuloManufacturadoService),

		ArticuloInsumoService:    container.ArticuloInsumoService,
		ArticuloInsumoController: controllers.NewArticuloInsumoController(container.ArticuloInsumoService),
	}
	return &app, nil
}

func (app *App) RegisterRoutes(router *gin.Engine) {

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	register := router.Group("/register")
	{
		register.POST("", app.LoginController.AddUsuario)
	}

	login := router.Group("/login")
	{
		login.POST("", app.LoginController.LoginUsuario)
	}

	usuarios := router.Group("/usuarios")
	{
		usuarios.GET("", app.LoginController.GetAllUsuarios)
		usuarios.GET("/:id", app.LoginController.GetUsuarioByID)
		usuarios.DELETE("/:id", app.LoginController.DeleteUsuarioByID)
		usuarios.PUT("", app.LoginController.UpdateUsuario)
	}

	instrumentoGroup := router.Group("/factura")
	{
		instrumentoGroup.GET("/:idFactura", app.FacturaController.GetByID)
		instrumentoGroup.POST("", app.FacturaController.AddFactura)
		instrumentoGroup.GET("/getAll", app.FacturaController.GetAll)
		instrumentoGroup.DELETE("/:idFactura", app.FacturaController.DeleteFactura)
		instrumentoGroup.PUT("", app.FacturaController.UpdateFactura)
	}

	productoGroup := router.Group("/pedido")
	{
		productoGroup.GET("/:idPedido", app.PedidoController.GetByID)
		productoGroup.POST("", app.PedidoController.AddPedido)
		productoGroup.GET("/getAll", app.PedidoController.GetAll)
		productoGroup.DELETE("/:idPedido", app.PedidoController.DeletePedido)
		productoGroup.PUT("", app.PedidoController.UpdatePedido)
	}

	articuloInsumo := router.Group("/articulo-insumo")
	{
		articuloInsumo.GET("/:id", app.ArticuloInsumoController.GetByID)
		articuloInsumo.POST("", app.ArticuloInsumoController.AddArticuloInsumo)
		articuloInsumo.GET("/getAll", app.ArticuloInsumoController.GetAll)
		articuloInsumo.DELETE("/:id", app.ArticuloInsumoController.DeleteArticuloInsumo)
		articuloInsumo.PUT("", app.ArticuloInsumoController.UpdateArticuloInsumo)
	}

	articuloManufacturadoDetalle := router.Group("/articulo-manufacturado-detalle")
	{
		articuloManufacturadoDetalle.GET("/:id", app.ArticuloManufacturadoDetalleController.GetByID)
		articuloManufacturadoDetalle.POST("", app.ArticuloManufacturadoDetalleController.AddArticuloManufacturadoDetalle)
		articuloManufacturadoDetalle.GET("/getAll", app.ArticuloManufacturadoDetalleController.GetAll)
		articuloManufacturadoDetalle.DELETE("/:id", app.ArticuloManufacturadoDetalleController.DeleteArticuloManufacturadoDetalle)
		articuloManufacturadoDetalle.PUT("", app.ArticuloManufacturadoDetalleController.UpdateArticuloManufacturadoDetalle)
	}

	articuloManufacturado := router.Group("/articulo-manufacturado")
	{
		articuloManufacturado.GET("/:id", app.ArticuloManufacturadoController.GetByID)
		articuloManufacturado.POST("", app.ArticuloManufacturadoController.AddArticuloManufacturado)
		articuloManufacturado.GET("/getAll", app.ArticuloManufacturadoController.GetAll)
		articuloManufacturado.DELETE("/:id", app.ArticuloManufacturadoController.DeleteArticuloManufacturado)
		articuloManufacturado.PUT("", app.ArticuloManufacturadoController.UpdateArticuloManufacturado)
	}

}

func (a *App) CerrarDB() {
	a.db.Close()
}
