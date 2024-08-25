package database

import (
	"api/src/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

type database struct {
	db *sql.DB
}

func NewDatabase() (DatabaseInterface, error) {
	db, err := OpenConnection()
	if err != nil {
		return nil, err
	}

	return &database{db}, nil
}

func (d *database) Query(query string, args ...any) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}

func (d *database) Close() {
	d.db.Close()
}

func (d *database) Prepare(query string) (*sql.Stmt, error) {
	return d.db.Prepare(query)
}

// OpenConnection open connection with database and return a connection
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDbConfig()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	db, err := sql.Open(conf.Drive, stringConnection)
	if err != nil {
		fmt.Print(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
