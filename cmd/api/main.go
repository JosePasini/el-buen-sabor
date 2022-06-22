package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	fmt.Println("Iniciando la app...")
	server, err := app.NewApp()
	server.RegisterRoutes(router)
	if err != nil {
		fmt.Println("Error al conectar la app.")
		server.CerrarDB()
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	err = router.Run(port)
	if err != nil {
		fmt.Println("Error al conectar la app en el puerto:", port)
		server.CerrarDB()
		return
	}
}
