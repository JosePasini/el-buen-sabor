package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type ILoginRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, usuario Usuario) error
	GetHash(ctx context.Context, tx *sqlx.Tx, usuario string) (string, error)
}

const (
	EMPLEADOS = 100
	COCINEROS = 200
	CLIENTES  = 300
)

type Usuario struct {
	ID       int     `json:"id"`
	Nombre   *string `json:"nombre"`
	Apellido *string `json:"apellido"`
	Usuario  *string `json:"usuario"`
	Mail     *string `json:"mail"`
	Hash     *string `json:"hash"`
	Rol      int     `json:"rol"`
}

func (u Usuario) isValid() bool {
	return u.ID > 0
}

func (u Usuario) GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
