package storage

import (
	"fmt"
	"database/sql"
	"note-service/config"
	_ "github.com/lib/pq"

)
func InitDB(cfg *config.Postgres) (*sql.DB, error){
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, err
}

func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}