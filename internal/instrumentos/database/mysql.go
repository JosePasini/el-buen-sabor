package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//MySQLConfig :: Struct configuración de MySQL
type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Database string
}

// GetConnectionString :: Obtiene string de conexion a db mysql
func (c MySQLConfig) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", c.User, c.Password, c.Host, c.Database)
}

// MySQL :: Struct de BD
type MySQL struct {
	*sqlx.DB
}

// NewMySQL :: conexion a DB mysql
func NewMySQL(config MySQLConfig) (*MySQL, error) {
	fmt.Println("[event:Connect] connecting to mysql DB...")
	db, err := sqlx.Open("mysql", config.GetConnectionString())
	if err != nil {
		fmt.Println("[event: fail_db_init][service: db_service] Could not start DB connection %s", err, err.Error())
		return nil, err
	}

	if err := mysql.SetLogger(log.New(os.Stderr, "mysql ", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)); err != nil {
		return nil, err
	}

	return &MySQL{db}, nil
}

// WithTransaction :: Control de transaction DB - Roolback and Commit
func (db MySQL) WithTransaction(ctx context.Context, f func(*sqlx.Tx) error) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			panic(p) // Volver a lanzar panic despues del roolback
		} else if err != nil {
			fmt.Println("error preventing transaction commit", err)
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = f(tx)
	return err
}

//Close :: Close db conn
func (db MySQL) Close() error {
	return db.DB.Close()
}
