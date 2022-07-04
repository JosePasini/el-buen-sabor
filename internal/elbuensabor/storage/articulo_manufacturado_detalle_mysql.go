package storage

type IArticuloManufacturadoDetalleRepository interface {
	// Insert(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	// GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.ArticuloManufacturadoDetalle, error)
	// GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.ArticuloManufacturadoDetalle, error)
	// Update(ctx context.Context, tx *sqlx.Tx, articulo_manufacturado_detalle domain.ArticuloManufacturadoDetalle) error
	// Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLArticuloManufacturadoDetalleRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

//  fecha (Date)
// número (int) montoDescuento (double) formaPago(String) nroTarjeta(String) totalVenta (double) totalCosto (double)

func NewMySQLArticuloManufacturadoDetalleRepository() *MySQLArticuloManufacturadoDetalleRepository {
	return &MySQLArticuloManufacturadoDetalleRepository{
		qInsert:     "INSERT INTO articulo_manufacturado_detalle (id) VALUES (?)",
		qGetByID:    "SELECT * FROM articulo_manufacturado_detalle WHERE id = ?",
		qGetAll:     "SELECT * FROM articulo_manufacturado_detalle",
		qDeleteById: "DELETE FROM articulo_manufacturado_detalle WHERE id = ?",
		qUpdate:     "UPDATE articulo_manufacturado_detalle SET fecha = COALESCE(?,fecha) WHERE id = ?",
	}
}
