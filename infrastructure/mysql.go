package infrastructure

import (
	"os"
	"strconv"

	"github.com/mkaiho/go-graphql-sample/adapter/mysql"
)

var _ mysql.Config = (*MySQLConfig)(nil)

type MySQLConfig struct{}

func (c *MySQLConfig) Host() string {
	return os.Getenv("DB_HOST")
}

func (c *MySQLConfig) Port() int {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	return port
}

func (c *MySQLConfig) DBName() string {
	return os.Getenv("DB_NAME")
}

func (c *MySQLConfig) User() string {
	return os.Getenv("DB_USER")
}

func (c *MySQLConfig) Password() string {
	return os.Getenv("DB_PASSWORD")
}
