package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		slog.Error("Error connecting to PostgreSQL", "err", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		slog.Error("Error pinging PostgreSQL", "err", err)
		return nil, err
	}

	slog.Info("Connected to PostgreSQL")

	// Criando o schema automaticamente
	schema := `
	CREATE TABLE IF NOT EXISTS exchange_rates (
		id SERIAL PRIMARY KEY,
		bid NUMERIC(10,4) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	_, err = db.Exec(schema)
	if err != nil {
		slog.Error("Error creating schema", "err", err)
		return nil, err
	}

	slog.Info("Database schema ensured")
	return db, nil
}
