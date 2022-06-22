package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/database"
	"github.com/JosePasiniMercadolibre/react-instrumentos/internal/instrumentos/domain"
	"github.com/jmoiron/sqlx"
)

type instrumentoDB struct {
	Id              int             `db:"id"`
	Instrumento     sql.NullString  `db:"instrumento"`
	Marca           sql.NullString  `db:"marca"`
	Modelo          sql.NullString  `db:"modelo"`
	Imagen          sql.NullString  `db:"imagen"`
	Precio          sql.NullFloat64 `db:"precio"`
	CostoEnvio      sql.NullFloat64 `db:"costo_envio"`
	CantidadVendida sql.NullInt32   `db:"cantidad_vendida"`
	Descripcion     sql.NullString  `db:"descripcion"`
}

func (i *instrumentoDB) toInstrumento() domain.Instrumento {
	return domain.Instrumento{
		Id:              i.Id,
		Instrumento:     database.ToStringP(i.Instrumento),
		Marca:           database.ToStringP(i.Marca),
		Modelo:          database.ToStringP(i.Modelo),
		Imagen:          database.ToStringP(i.Imagen),
		Precio:          database.ToFloat64P(i.Precio),
		CostoEnvio:      database.ToFloat64P(i.CostoEnvio),
		CantidadVendida: database.ToIntP(i.CantidadVendida),
		Descripcion:     database.ToStringP(i.Descripcion),
	}
}

type MySQLInstrumentoRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

func NewMySQLInstrumentoRepository() *MySQLInstrumentoRepository {
	return &MySQLInstrumentoRepository{
		qInsert:     "INSERT INTO instrumentos (instrumento, marca, modelo, imagen, precio, costo_envio, cantidad_vendida, descripcion) VALUES (?,?,?,?,?,?,?,?);",
		qGetByID:    "SELECT id, instrumento, marca, modelo, imagen, precio, costo_envio, cantidad_vendida, descripcion FROM instrumentos WHERE id = ?",
		qGetAll:     "SELECT id, instrumento, marca, modelo, imagen, precio, costo_envio, cantidad_vendida, descripcion FROM instrumentos",
		qDeleteById: "DELETE FROM instrumentos WHERE id = ?",
		qUpdate:     "UPDATE instrumentos SET instrumento = COALESCE(?,instrumento), marca = COALESCE(?,marca) , modelo = COALESCE(?,modelo), imagen = COALESCE(?,imagen), precio = COALESCE(?,precio), costo_envio = COALESCE(?,costo_envio), cantidad_vendida = COALESCE(?,cantidad_vendida), descripcion = COALESCE(?,descripcion) WHERE id = ?",
	}
}

func (i *MySQLInstrumentoRepository) Update(ctx context.Context, tx *sqlx.Tx, ins domain.Instrumento) error {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, ins.Instrumento, ins.Marca, ins.Modelo, ins.Imagen, ins.Precio, ins.CostoEnvio, ins.CantidadVendida, ins.Descripcion, ins.Id)
	return err
}

func (i *MySQLInstrumentoRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := i.qDeleteById
	_, err := tx.ExecContext(ctx, query, id)
	return err
}

func (i *MySQLInstrumentoRepository) Insert(ctx context.Context, tx *sqlx.Tx, ins domain.Instrumento) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, ins.Instrumento, ins.Marca, ins.Modelo, ins.Imagen, ins.Precio, ins.CostoEnvio, ins.CantidadVendida, ins.Descripcion)
	return err
}

func (i *MySQLInstrumentoRepository) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*domain.Instrumento, error) {
	query := i.qGetByID
	var instrumento instrumentoDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&instrumento)
	if err != nil {
		return nil, err
	}
	inst := instrumento.toInstrumento()
	return &inst, nil
}

func (i *MySQLInstrumentoRepository) GetAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Instrumento, error) {
	query := i.qGetAll
	instruments := make([]domain.Instrumento, 0)

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var instrumento instrumentoDB
		if err := rows.StructScan(&instrumento); err != nil {
			return instruments, err
		}
		instruments = append(instruments, instrumento.toInstrumento())
	}
	return instruments, nil
}
