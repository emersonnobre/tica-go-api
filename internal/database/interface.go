package database

import "database/sql"

type Database interface {
	Connect() (*sql.DB, error)
	Query(query string) (any, error)
}

type DatabaseConfig struct {
	Host         string
	DatabaseName string
	User         string
	Password     string
	Tls          string
}
