package database

import (
	"database/sql"
)

type DatabaseInterface interface {
	Query(string, ...any) (*sql.Rows, error)
	Prepare(string) (*sql.Stmt, error)
	Close()
}
