package services

import (
	"context"
	"fmt"

	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/database"
	"github.com/JosePasiniMercadolibre/el-buen-sabor/internal/elbuensabor/domain"
	"github.com/jmoiron/sqlx"
)

type ILoginService interface {
	AddUsuario(context.Context, domain.Usuario) error
	LoginUsuario(context.Context, domain.Login) (bool, error)
	GetAllUsuarios(context.Context) ([]domain.UsuarioResponse, error)
	GetUsuarioByID(context.Context, int) (domain.UsuarioResponse, error)
	GetUsuarioByEmail(context.Context, string) (domain.UsuarioResponse, error)
	DeleteUsuarioByID(context.Context, int) (bool, error)
	UpdateUsuario(context.Context, domain.Usuario) (domain.Usuario, error)
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
	fmt.Println("User SERVICE:", &usuario)
	bytesReturned, err := usuario.GeneratePassword(*usuario.Hash)
	if err != nil {
		return err
	}

	pass := string(bytesReturned)
	usuario.Hash = &pass
	fmt.Println("User Hash:", usuario.Hash)
	fmt.Println("User Hash:", *usuario.Hash)
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		err = s.repository.Insert(ctx, tx, usuario)
		return err
	})
	return err
}

func (s *LoginService) LoginUsuario(ctx context.Context, usuario domain.Login) (bool, error) {
	var bool = false
	var hashEmailDB string
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		hashEmailDB, err = s.repository.GetHash(ctx, tx, usuario.Email)
		return err
	})
	fmt.Println("Hash User Email:", hashEmailDB)
	bool, err = domain.ValidatePassword(usuario.Hash, hashEmailDB)
	return bool, nil
}
func (s *LoginService) GetAllUsuarios(ctx context.Context) ([]domain.UsuarioResponse, error) {
	var listaUsuarios []domain.UsuarioResponse
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		listaUsuarios, err = s.repository.GetAllUsuarios(ctx, tx)
		return err
	})
	fmt.Println("Lista Usuarios:", listaUsuarios)
	return listaUsuarios, nil
}

func (s *LoginService) GetUsuarioByID(ctx context.Context, id int) (domain.UsuarioResponse, error) {
	var usuario domain.UsuarioResponse
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		usuario, err = s.repository.GetUsuarioByID(ctx, tx, id)
		return err
	})
	if err != nil {
		return domain.UsuarioResponse{}, err
	}
	fmt.Println("Usuario:", usuario)
	return usuario, nil
}

func (s *LoginService) GetUsuarioByEmail(ctx context.Context, email string) (domain.UsuarioResponse, error) {
	var usuario domain.UsuarioResponse
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		usuario, err = s.repository.GetUsuarioByEmail(ctx, tx, email)
		return err
	})
	if err != nil {
		return domain.UsuarioResponse{}, err
	}
	fmt.Println("Usuario:", usuario)
	return usuario, nil
}

func (s *LoginService) DeleteUsuarioByID(ctx context.Context, id int) (bool, error) {
	var ok bool
	var err error
	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		ok, err = s.repository.DeleteUsuarioByID(ctx, tx, id)
		return err
	})
	fmt.Println("Eliminado?:", ok)
	return ok, nil
}

func (s *LoginService) UpdateUsuario(ctx context.Context, usuario domain.Usuario) (domain.Usuario, error) {
	var usuarioResponse domain.Usuario
	var err error

	bytesReturned, err := usuario.GeneratePassword(*usuario.Hash)
	if err != nil {
		return domain.Usuario{}, err
	}

	pass := string(bytesReturned)
	usuario.Hash = &pass

	err = s.db.WithTransaction(ctx, func(tx *sqlx.Tx) error {
		usuarioResponse, err = s.repository.UpdateUsuario(ctx, tx, usuario)
		return err
	})
	//fmt.Println("Updateado?:", usuarioResponse)
	if err != nil {
		return domain.Usuario{}, err
	}
	fmt.Println("Si, updateado:", usuarioResponse)
	return usuarioResponse, nil
}
