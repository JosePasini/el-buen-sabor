package elbuensabor

import "github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"

type AppConfig struct {
	DB database.MySQLConfig
}
