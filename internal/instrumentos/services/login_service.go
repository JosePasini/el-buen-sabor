package services

import (
	"context"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/instrumentos/domain"
	"github.com/jmoiron/sqlx"
)

type ILoginService interface {
	AddUsuario(context.Context, domain.Usuario) error
}

type LoginService struct {
	db         database.DB
	repository domain.ILoginRepository
}

func NewLoginService(db database.DB, repository domain.ILoginRepository) *LoginService {
	return &LoginService{db, repository}
}

func (s *LoginService) AddUsuario(ctx context.Context, usuario domain.Usuario) error {
	var err error
	fmt.Println("User SERVICE:")
	fmt.Println("User SERVICE:", &usuario)
	bytesReturned, err := usuario.GeneratePassword(*usuario.Hash)
	if err != nil {
		return err
	}

	fmt.Println("Bytes:", bytesReturned)
	pass := string(bytesReturned)
	usuario.Hash = &pass
	fmt.Println("Pass:", pass)
	fmt.Println("User Hash:", usuario.Hash)
	fmt.Println("User Hash:", *usuario.Hash)
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, usuario)
		return err
	})
	return err
}
