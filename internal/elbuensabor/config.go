package instrumentos

import "github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/database"

type AppConfig struct {
	DB database.MySQLConfig
}
