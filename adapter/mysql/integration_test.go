//go:build mysql

package mysql

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
)

var testDB *sqlx.DB
var _ Config = (*testConfig)(nil)

type testConfig struct{}

func (c *testConfig) Host() string {
	value := os.Getenv("DB_HOST")
	if len(value) == 0 {
		value = "mysqldb"
	}
	return value
}
func (c *testConfig) Port() int {
	value, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		value = 3306
	}
	return value
}
func (c *testConfig) DBName() string {
	value := os.Getenv("DB_NAME")
	if len(value) == 0 {
		value = "devdb"
	}
	return value
}
func (c *testConfig) User() string {
	value := os.Getenv("DB_USER")
	if len(value) == 0 {
		value = "devuser"
	}
	return value
}
func (c *testConfig) Password() string {
	value := os.Getenv("DB_PASSWORD")
	if len(value) == 0 {
		value = "devdev"
	}
	return value
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	var err error
	testDB, err = NewDB(&testConfig{})
	if err != nil {
		panic(err)
	}
}

func load(fileName string) error {
	path, err := filepath.Abs(fmt.Sprintf("./fixture/%s", fileName))
	if err != nil {
		return err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for _, q := range strings.Split(string(b), ";") {
		if len(strings.TrimSpace(q)) != 0 {
			_, err = testDB.Exec(q)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
