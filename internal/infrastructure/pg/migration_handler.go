package pg

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"
)

func initializeDatabase(db *sql.DB) {
	createSchemaAndTable(db)
	runMigrations(db)
}

// createSchemaAndTable cria o esquema sistema e a tabela migration se não existirem
func createSchemaAndTable(db *sql.DB) {
	schemaSQL := `
    CREATE SCHEMA IF NOT EXISTS sistema;
    CREATE TABLE IF NOT EXISTS sistema.migration (
        id SERIAL PRIMARY KEY,
        filename VARCHAR(255) NOT NULL,
        executed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(schemaSQL)
	if err != nil {
		log.Fatalf("Erro ao criar esquema e tabela: %v", err)
	}
}

// runMigrations executa as migrações pendentes
func runMigrations(db *sql.DB) {
	migrationDir := "db/migrations"

	files, err := os.ReadDir(migrationDir)
	if err != nil {
		log.Fatalf("Erro ao ler diretório de migrações: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			if !isMigrationExecuted(db, file.Name()) {
				executeMigration(db, filepath.Join(migrationDir, file.Name()), file.Name())
			}
		}
	}
}

// isMigrationExecuted verifica se a migração já foi executada
func isMigrationExecuted(db *sql.DB, filename string) bool {
	var count int
	query := `SELECT COUNT(*) FROM sistema.migration WHERE filename = $1`
	err := db.QueryRow(query, filename).Scan(&count)
	if err != nil {
		log.Fatalf("Erro ao verificar migração: %v", err)
	}
	return count > 0
}

// executeMigration executa a migração e registra na tabela
func executeMigration(db *sql.DB, filepath string, filename string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo de migração: %v", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatalf("Erro ao executar migração: %v", err)
	}

	query := `INSERT INTO sistema.migration (filename, executed_at) VALUES ($1, $2)`
	_, err = db.Exec(query, filename, time.Now())
	if err != nil {
		log.Fatalf("Erro ao registrar migração: %v", err)
	}
}
