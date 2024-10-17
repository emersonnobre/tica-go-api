package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatabase struct {
	config DatabaseConfig
}

func NewMySQLDatabase() MySQLDatabase {
	return MySQLDatabase{
		config: DatabaseConfig{
			Host:         os.Getenv("MYSQL_HOST"),
			DatabaseName: os.Getenv("MYSQL_DATABASE"),
			User:         os.Getenv("MYSQL_USER"),
			Password:     os.Getenv("MYSQL_PASSWORD"),
			Tls:          os.Getenv("MYSQL_TLS"),
		},
	}
}

func (m *MySQLDatabase) Connect() (*sql.DB, error) {
	connection, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=%s",
			m.config.User,
			m.config.Password,
			m.config.Host,
			m.config.DatabaseName,
			m.config.Tls))
	return connection, err
}
