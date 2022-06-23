package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type ILoginRepository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, usuario Usuario) error
}

type Usuario struct {
	ID       int     `json:"id"`
	Nombre   *string `json:"nombre"`
	Apellido *string `json:"apellido"`
	Usuario  *string `json:"usuario"`
	Mail     *string `json:"mail"`
	Hash     *string `json:"hash"`
	//HashedPassword []byte `json:"-"`
}

func (u Usuario) isValid() bool {
	return u.ID > 0
}

func (u Usuario) GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
