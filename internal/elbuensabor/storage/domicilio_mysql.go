package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type domicilioDB struct {
	ID        int            `json:"id" db:"id"`
	Calle     sql.NullString `json:"calle" db:"calle"`
	Numero    sql.NullInt32  `json:"numero" db:"numero"`
	Localidad sql.NullString `json:"localidad" db:"localidad"`
	IDUsuario sql.NullInt32  `json:"id_usuario" db:"id_usuario"`
}

func (a *domicilioDB) toDomicilio() domain.Domicilio {
	return domain.Domicilio{
		ID:        a.ID,
		Calle:     database.ToStringP(a.Calle),
		Numero:    database.ToIntP(a.Numero),
		Localidad: database.ToStringP(a.Localidad),
		IDUsuario: database.ToIntP(a.IDUsuario),
	}
}

type IDomicilioRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, domicilio domain.Domicilio) error
	//GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Domicilio, error)
	Update(ctx context.Context, tx *sqlx.Tx, domicilio domain.Domicilio) error
	GetAllDomicilioByUsuario(ctx context.Context, tx *sqlx.Tx, idUsuario int) ([]domain.Domicilio, error)
	// Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type MySQLDomicilioRepository struct {
	qInsert                   string
	qUpdate                   string
	qGetAllDomicilioByUsuario string
	// qGetByID    string
	// qDeleteById string
}

func NewMySQLDomicilioRepository() *MySQLDomicilioRepository {
	return &MySQLDomicilioRepository{
		qInsert:                   "INSERT INTO domicilio (calle, numero, localidad, id_usuario) VALUES (?,?,?,?)",
		qUpdate:                   "UPDATE domicilio SET calle = COALESCE(?,calle), numero = COALESCE(?,numero),localidad = COALESCE(?,localidad) WHERE id = ?",
		qGetAllDomicilioByUsuario: "SELECT d.id, d.calle, d.numero, d.localidad, d.id_domicilio FROM domicilio d JOIN usuarios u ON d.id_usuario = u.id WHERE d.id_usuario = ?",
		//qGetByID:    "SELECT id, cantidad, id_articulo_manufacturado, id_articulo_insumo FROM articulo_manufacturado_detalle WHERE id = ?",
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
	_, err := tx.ExecContext(ctx, query, domicilio.Calle, domicilio.Numero, domicilio.Localidad, domicilio.IDUsuario)
	return err
}
func (i *MySQLDomicilioRepository) Update(ctx context.Context, tx *sqlx.Tx, domicilio domain.Domicilio) error {
	fmt.Println("domicilio:", domicilio)
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, domicilio.Calle, domicilio.Numero, domicilio.Localidad, domicilio.ID)
	return err
}

func (i *MySQLDomicilioRepository) GetAllDomicilioByUsuario(ctx context.Context, tx *sqlx.Tx, idUsuario int) ([]domain.Domicilio, error) {
	fmt.Println("idUsuario:", idUsuario)
	query := i.qGetAllDomicilioByUsuario
	domicilios := make([]domain.Domicilio, 0)

	rows, err := tx.QueryxContext(ctx, query, idUsuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dom domicilioDB
		if err := rows.StructScan(&dom); err != nil {
			return domicilios, err
		}
		domicilios = append(domicilios, dom.toDomicilio())
	}
	return domicilios, nil

}
