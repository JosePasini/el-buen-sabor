package app

import (
	"net/http"

	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/controllers"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/database"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	db     database.DB
	Config instrumentos.AppConfig

	InstrumentoService    services.IInstrumentoService
	InstrumentoController controllers.IInstrumentoController
}

func NewApp() (*App, error) {
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
		Config:                config,
		InstrumentoService:    container.InstrumentoService,
		InstrumentoController: controllers.NewInstrumentoController(container.InstrumentoService),
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
			"message": "Hello World",
		})
	})

	instrumentoGroup := router.Group("/instrumento")
	{
		instrumentoGroup.GET("/:idInstrumento", app.InstrumentoController.GetByID)
		instrumentoGroup.POST("", app.InstrumentoController.AddInstrument)
		instrumentoGroup.GET("/getAll", app.InstrumentoController.GetAll)
		instrumentoGroup.DELETE("/:idInstrumento", app.InstrumentoController.DeleteInstrument)
		instrumentoGroup.PUT("", app.InstrumentoController.UpdateInstrument)
	}

}

func (a *App) CerrarDB() {
	a.db.Close()
}
