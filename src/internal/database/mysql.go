package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatabase struct {
	Config DatabaseConfig
}

func NewMySQLDatabase() MySQLDatabase {
	return MySQLDatabase{
		Config: DatabaseConfig{
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
		fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=%s&parseTime=True",
			m.Config.User,
			m.Config.Password,
			m.Config.Host,
			m.Config.DatabaseName,
			m.Config.Tls))
	return connection, err
}
