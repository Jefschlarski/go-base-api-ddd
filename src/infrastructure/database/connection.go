package database

import (
	"api/src/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

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
