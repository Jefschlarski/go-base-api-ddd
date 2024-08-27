package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"taskmanager/internal/configs"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

var (
	db   *sql.DB
	once sync.Once
)

// GetPostgresDB retorna a instância singleton da conexão com o banco de dados postgres.
func GetPostgresDB() *sql.DB {
	once.Do(func() {
		db = openPostgresConnection()
	})
	return db
}

// openPostgresConnection open connection with postgres database and return a connection
func openPostgresConnection() *sql.DB {

	conf := configs.GetDbConfig()

	db, err := sql.Open(conf.Drive, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database))

	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Erro ao se conectar ao banco de dados: %v", err)
	}

	return db
}
