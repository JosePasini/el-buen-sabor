package storage

type IArticuloInsumoRepository interface {
	// Insert(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	// GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturadoDetalle, error)
	// GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturadoDetalle, error)
	// Update(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	// Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLArticuloInsumoRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

//  fecha (Date)
// número (int) montoDescuento (double) formaPago(String) nroTarjeta(String) totalVenta (double) totalCosto (double)

func NewMySQLArticuloInsumoRepository() *MySQLArticuloInsumoRepository {
	return &MySQLArticuloInsumoRepository{
		qInsert:     "INSERT INTO articulo_insumo (id) VALUES (?)",
		qGetByID:    "SELECT * FROM articulo_insumo WHERE id = ?",
		qGetAll:     "SELECT * FROM articulo_insumo",
		qDeleteById: "DELETE FROM articulo_insumo WHERE id = ?",
		qUpdate:     "UPDATE articulo_insumo SET fecha = COALESCE(?,fecha) WHERE id = ?",
	}
}
