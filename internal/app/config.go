package app

import (
	"strings"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/database"
)

// NewConfig :: Carga de configuración inicial
func NewConfig(scope string) (instrumentos.AppConfig, error) {
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
			User:     "root_produ",
			Password: "",
			Host:     "localhost",
			Database: "react",
		},
	}, nil
}
