package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/joho/godotenv"
)

// NewConfig :: Carga de configuración inicial
func NewConfig(scope string) (elbuensabor.AppConfig, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No se pudo cargar el archivo .env")
	}

	USER_ENV := os.Getenv("DB_USERNAME")
	PASS_ENV := os.Getenv("DB_PASS")
	HOST_ENV := os.Getenv("DB_HOST")
	NAME_ENV := os.Getenv("DB_NAME")

	fmt.Println(":::::", HOST_ENV, " ", USER_ENV, " ", PASS_ENV, " ", NAME_ENV)

	if !strings.Contains(scope, "prod") {
		return elbuensabor.AppConfig{
			DB: database.MySQLConfig{
				User:     "root",
				Password: "",
				Host:     "localhost",
				Database: "elbuensabor",
			},
		}, nil
	}

	return elbuensabor.AppConfig{
		DB: database.MySQLConfig{
			User:     USER_ENV,
			Password: PASS_ENV,
			Host:     HOST_ENV,
			Database: NAME_ENV,
		},
	}, nil
}
