package storage

type IArticuloManufacturadoRepository interface {
	// Insert(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	// GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturadoDetalle, error)
	// GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturadoDetalle, error)
	// Update(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	// Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLArticuloManufacturadoRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

//  fecha (Date)
// número (int) montoDescuento (double) formaPago(String) nroTarjeta(String) totalVenta (double) totalCosto (double)

func NewMySQLArticuloManufacturadoRepository() *MySQLArticuloManufacturadoRepository {
	return &MySQLArticuloManufacturadoRepository{
		qInsert:     "INSERT INTO articulo_manufacturado (id) VALUES (?)",
		qGetByID:    "SELECT * FROM articulo_manufacturado WHERE id = ?",
		qGetAll:     "SELECT * FROM articulo_manufacturado",
		qDeleteById: "DELETE FROM articulo_manufacturado WHERE id = ?",
		qUpdate:     "UPDATE articulo_manufacturado SET fecha = COALESCE(?,fecha) WHERE id = ?",
	}
}
