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

	godotenv.Load()
	// Variables SCOPE LOCAL
	HOST_LOCAL := os.Getenv("HOST_LOCAL")
	USERNAME_LOCAL := os.Getenv("USERNAME_LOCAL")
	DB_LOCAL := os.Getenv("DB_LOCAL")
	PASSWORD_LOCAL := os.Getenv("PASSWORD_LOCAL")

	// Variables SCOPE PRODUCTIVO
	USER_ENV := os.Getenv("DB_USERNAME")
	PASS_ENV := os.Getenv("DB_PASS")
	HOST_ENV := os.Getenv("DB_HOST")
	NAME_ENV := os.Getenv("DB_NAME")
	fmt.Println("BD", os.Getenv("DB_HOST"))

	if !strings.Contains(scope, "prod") {
		fmt.Println(":: connecting LOCAL.... :::", USERNAME_LOCAL, " :: ", PASSWORD_LOCAL, " :: ", HOST_LOCAL, " :: ", DB_LOCAL)
	} else {
		fmt.Println(":: connecting PROD... :::", HOST_ENV, " ", USER_ENV, " ", PASS_ENV, " ", NAME_ENV)
	}

	if !strings.Contains(scope, "prod") {
		return elbuensabor.AppConfig{
			DB: database.MySQLConfig{
				User:     USERNAME_LOCAL,
				Password: PASSWORD_LOCAL,
				Host:     HOST_LOCAL,
				Database: DB_LOCAL,
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
