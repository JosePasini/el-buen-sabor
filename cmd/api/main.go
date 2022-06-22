package main

import (
	"fmt"
	"net/http"
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

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	router.Run(":" + port)

	// fmt.Println("Hello World")
	// router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.1.2"})

	// fmt.Println("Iniciando la app...")
	// server, err := app.NewApp()
	// server.RegisterRoutes(router)
	// if err != nil {
	// 	fmt.Println("Error al conectar la app.")
	// 	server.CerrarDB()
	// 	return
	// }

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = ":8080"
	// }

	// err = router.Run(port)
	// if err != nil {
	// 	fmt.Println("Error al conectar la app en el puerto:", port)
	// 	server.CerrarDB()
	// 	return
	// }
}
