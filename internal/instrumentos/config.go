package instrumentos

import "github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/database"

type AppConfig struct {
	DB database.MySQLConfig
}
