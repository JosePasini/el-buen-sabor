package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/domain"
	"github.com/jmoiron/sqlx"
)

type loginDB struct {
	ID       int            `db:"id"`
	Nombre   sql.NullString `db:"nombre"`
	Apellido sql.NullString `db:"apellido"`
	Mail     sql.NullString `db:"mail"`
	Usuario  sql.NullString `db:"usuario"`
	Hash     sql.NullString `db:"hash"`
}

func (i *loginDB) toLoginDB() domain.Usuario {
	return domain.Usuario{
		ID:       i.ID,
		Nombre:   database.ToStringP(i.Nombre),
		Apellido: database.ToStringP(i.Apellido),
		Mail:     database.ToStringP(i.Mail),
		Usuario:  database.ToStringP(i.Usuario),
		Hash:     database.ToStringP(i.Hash),
	}
}

type MySQLLoginRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
}

func NewMySQLLoginRepository() *MySQLLoginRepository {
	return &MySQLLoginRepository{
		qInsert:     "INSERT INTO usuarios (nombre, apellido, mail, usuario, hash) VALUES (?,?,?,?,?);",
		qGetAll:     "SELECT * FROM usuarios",
		qGetByID:    "SELECT * FROM usuarios WHERE id = ?",
		qDeleteById: "DELETE FROM usuarios WHERE id = ?",
		qUpdate:     "UPDATE usuarios SET nombre = COALESCE(?,nombre), apellido = COALESCE(?,apellido) , mail = COALESCE(?,mail), hash = COALESCE(?,hash) WHERE id = ?",
	}
}

func (i *MySQLLoginRepository) Insert(ctx context.Context, tx *sqlx.Tx, user domain.Usuario) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, user.Nombre, user.Apellido, user.Mail, user.Usuario, user.Hash)
	return err
}
