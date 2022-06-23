package main

import (
	"fmt"
	"os"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies([]string{"192.168.1.2"})
	fmt.Println("Iniciando la app...")
	server, err := app.NewApp()
	server.RegisterRoutes(router)
	if err != nil {
		fmt.Println("Error al conectar la app.")
		server.CerrarDB()
		return
	}
	router.Run(":" + port)
}
