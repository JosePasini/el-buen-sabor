package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/database"
	"github.com/joho/godotenv"
)

// NewConfig :: Carga de configuración inicial
func NewConfig(scope string) (instrumentos.AppConfig, error) {

	godotenv.Load()

	USER_ENV := os.Getenv("DB_USER")
	PASS_ENV := os.Getenv("DB_PASS")
	HOST_ENV := os.Getenv("DB_HOST")
	NAME_ENV := os.Getenv("DB_NAME")

	fmt.Println(":::::", HOST_ENV, " ", USER_ENV, " ", PASS_ENV, " ", NAME_ENV)

	if !strings.Contains(scope, "prod") {
		return instrumentos.AppConfig{
			DB: database.MySQLConfig{
				User:     "root",
				Password: "",
				Host:     "localhost",
				Database: "react",
			},
		}, nil
	}

	return instrumentos.AppConfig{
		DB: database.MySQLConfig{
			User:     USER_ENV,
			Password: PASS_ENV,
			Host:     HOST_ENV,
			Database: NAME_ENV,
		},
	}, nil
	// return instrumentos.AppConfig{
	// 	DB: database.MySQLConfig{
	// 		User:     "baac4c3d3bb29e",
	// 		Password: "408399dd",
	// 		Host:     "us-cdbr-east-05.cleardb.net",
	// 		Database: "heroku_9952cf2f0b46460",
	// 	},
	// }, nil
}
