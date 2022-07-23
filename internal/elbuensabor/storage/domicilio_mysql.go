package storage

import (
	"context"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type IDomicilioRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, domicilio domain.Domicilio) error
	//GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Domicilio, error)
	Update(ctx context.Context, tx *sqlx.Tx, domicilio domain.Domicilio) error
	// Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLDomicilioRepository struct {
	qInsert string
	qUpdate string
	// qGetByID    string
	// qGetAll     string
	// qDeleteById string
}

func NewMySQLDomicilioRepository() *MySQLDomicilioRepository {
	return &MySQLDomicilioRepository{
		qInsert: "INSERT INTO domicilio (calle, numero, localidad) VALUES (?,?,?)",
		qUpdate: "UPDATE domicilio SET calle = COALESCE(?,calle), numero = COALESCE(?,numero),localidad = COALESCE(?,localidad) WHERE id = ?",
		//qGetByID:    "SELECT id, cantidad, id_articulo_manufacturado, id_articulo_insumo FROM articulo_manufacturado_detalle WHERE id = ?",
		//qGetAll:     "SELECT id, cantidad, id_articulo_manufacturado, id_articulo_insumo FROM articulo_manufacturado_detalle",
		//qDeleteById: "DELETE FROM articulo_manufacturado_detalle WHERE id = ?",
	}
}

// func (i *MySQLArticuloManufacturadoDetalleRepository) Update(ctx context.Context, tx *sqlx.Tx, artManu domain.ArticuloManufacturadoDetalle) error {
// 	query := i.qUpdate
// 	_, err := tx.ExecContext(ctx, query, artManu.Cantidad, artManu.IDArticuloManufacturado, artManu.IDArticuloInsumo, artManu.ID)
// 	return err
// }

// func (i *MySQLArticuloManufacturadoDetalleRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
// 	query := i.qDeleteById
// 	_, err := tx.ExecContext(ctx, query, id)
// 	return err
// }

func (i *MySQLDomicilioRepository) Insert(ctx context.Context, tx *sqlx.Tx, domicilio domain.Domicilio) error {
	fmt.Println("domicilio:", domicilio)
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, domicilio.Calle, domicilio.Numero, domicilio.Localidad)
	return err
}
func (i *MySQLDomicilioRepository) Update(ctx context.Context, tx *sqlx.Tx, domicilio domain.Domicilio) error {
	fmt.Println("domicilio:", domicilio)
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, domicilio.Calle, domicilio.Numero, domicilio.Localidad, domicilio.ID)
	return err
}
