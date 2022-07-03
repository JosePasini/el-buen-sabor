package storage

import (
	"context"
	"database/sql"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type loginDB struct {
	ID       int            `db:"id"`
	Nombre   sql.NullString `db:"nombre"`
	Apellido sql.NullString `db:"apellido"`
	Email    sql.NullString `db:"email"`
	Usuario  sql.NullString `db:"usuario"`
	Hash     sql.NullString `db:"hash"`
	Rol      int            `db:"rol"`
}

type loginResponseDB struct {
	ID       int            `db:"id"`
	Nombre   sql.NullString `db:"nombre"`
	Apellido sql.NullString `db:"apellido"`
	Email    sql.NullString `db:"email"`
	Usuario  sql.NullString `db:"usuario"`
	Rol      int            `db:"rol"`
}

func (i *loginResponseDB) toLoginResponseDB() domain.UsuarioResponse {
	return domain.UsuarioResponse{
		ID:       i.ID,
		Nombre:   database.ToStringP(i.Nombre),
		Apellido: database.ToStringP(i.Apellido),
		Email:    database.ToStringP(i.Email),
		Usuario:  database.ToStringP(i.Usuario),
		Rol:      i.Rol,
	}
}

func (i *loginDB) toLoginDB() domain.Usuario {
	return domain.Usuario{
		ID:       i.ID,
		Nombre:   database.ToStringP(i.Nombre),
		Apellido: database.ToStringP(i.Apellido),
		Email:    database.ToStringP(i.Email),
		Usuario:  database.ToStringP(i.Usuario),
		Hash:     database.ToStringP(i.Hash),
		Rol:      i.Rol,
	}
}

type MySQLLoginRepository struct {
	qInsert     string
	qGetByID    string
	qGetAll     string
	qDeleteById string
	qUpdate     string
	qGetHash    string
	qGetByEmail string
}

func NewMySQLLoginRepository() *MySQLLoginRepository {
	return &MySQLLoginRepository{
		qInsert:     "INSERT INTO usuarios (nombre, apellido, email, usuario, hash) VALUES (?,?,?,?,?);",
		qGetAll:     "SELECT id, nombre, apellido, email, usuario, rol FROM usuarios",
		qGetByID:    "SELECT id, nombre, apellido, email, usuario, rol FROM usuarios WHERE id = ?",
		qGetByEmail: "SELECT id, nombre, apellido, email, usuario, rol FROM usuarios WHERE email = ?",
		qDeleteById: "DELETE FROM usuarios WHERE id = ?",
		qUpdate:     "UPDATE usuarios SET nombre = COALESCE(?,nombre), apellido = COALESCE(?,apellido), usuario = COALESCE(?,usuario) , email = COALESCE(?,email), hash = COALESCE(?,hash) WHERE id = ?",
		qGetHash:    "SELECT hash FROM usuarios WHERE email = ?",
	}
}

func (i *MySQLLoginRepository) Insert(ctx context.Context, tx *sqlx.Tx, user domain.Usuario) error {
	query := i.qInsert
	_, err := tx.ExecContext(ctx, query, user.Nombre, user.Apellido, user.Email, user.Usuario, user.Hash)
	return err
}

func (i *MySQLLoginRepository) GetHash(ctx context.Context, tx *sqlx.Tx, email string) (string, error) {
	query := i.qGetHash
	var hashUser string

	row := tx.QueryRowxContext(ctx, query, email)
	err := row.Scan(&hashUser)
	if err != nil {
		return hashUser, err
	}
	return hashUser, nil
}

func (i *MySQLLoginRepository) GetAllUsuarios(ctx context.Context, tx *sqlx.Tx) ([]domain.UsuarioResponse, error) {
	query := i.qGetAll
	var listaUsuarios []domain.UsuarioResponse
	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var listaUsuariosDB loginResponseDB
		if err := rows.StructScan(&listaUsuariosDB); err != nil {
			return listaUsuarios, err
		}
		listaUsuarios = append(listaUsuarios, listaUsuariosDB.toLoginResponseDB())
	}
	return listaUsuarios, nil
}

func (i *MySQLLoginRepository) GetUsuarioByID(ctx context.Context, tx *sqlx.Tx, id int) (domain.UsuarioResponse, error) {
	query := i.qGetByID
	var usuario loginResponseDB

	row := tx.QueryRowxContext(ctx, query, id)
	err := row.StructScan(&usuario)
	if err != nil {
		return domain.UsuarioResponse{}, err
	}
	usuarioResponse := usuario.toLoginResponseDB()
	return usuarioResponse, nil
}

func (i *MySQLLoginRepository) GetUsuarioByEmail(ctx context.Context, tx *sqlx.Tx, email string) (domain.UsuarioResponse, error) {
	query := i.qGetByEmail
	var usuario loginResponseDB

	row := tx.QueryRowxContext(ctx, query, email)
	err := row.StructScan(&usuario)
	if err != nil {
		return domain.UsuarioResponse{}, err
	}
	usuarioResponse := usuario.toLoginResponseDB()
	return usuarioResponse, nil
}

func (i *MySQLLoginRepository) DeleteUsuarioByID(ctx context.Context, tx *sqlx.Tx, id int) (bool, error) {
	query := i.qDeleteById

	sqlResult, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := sqlResult.RowsAffected()
	if rowsAffected > 0 {
		return true, nil
	}
	return false, nil
}

func (i *MySQLLoginRepository) UpdateUsuario(ctx context.Context, tx *sqlx.Tx, u domain.Usuario) (domain.Usuario, error) {
	query := i.qUpdate
	_, err := tx.ExecContext(ctx, query, u.Nombre, u.Apellido, u.Usuario, u.Email, u.Hash, u.ID)
	if err != nil {
		return u, err
	}
	return domain.Usuario{}, err
}
