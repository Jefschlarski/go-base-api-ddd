package database

import (
	"api/src/common/errors"
	"api/src/configs"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

// OpenConnection open connection with database and return a connection
func OpenConnection() (*sql.DB, *errors.Error) {
	conf := configs.GetDbConfig()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	db, err := sql.Open(conf.Drive, stringConnection)
	if err != nil {
		fmt.Print(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	return db, nil
}
