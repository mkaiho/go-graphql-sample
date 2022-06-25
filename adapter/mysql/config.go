package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config interface {
	Host() string
	Port() int
	DBName() string
	User() string
	Password() string
}

func NewDB(config Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=UTC", config.User(), config.Password(), config.Host(), config.DBName())
	return sqlx.Open("mysql", dsn)
}
