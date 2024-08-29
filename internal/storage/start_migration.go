package storage

import (
    "database/sql"
    "fmt"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "note-service/config"
)

type MigrationService struct {
    DB  *sql.DB
    Cfg config.Postgres
}


func NewMigrationService(db *sql.DB, cfg config.Postgres) *MigrationService {
    return &MigrationService{DB: db, Cfg: cfg}
}

func (ms *MigrationService) Run() error {
    driver, err := postgres.WithInstance(ms.DB, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("failed to create migration driver: %w", err)
    }

    // Construct and log the path to the migrations directory
    migrationsPath := "file://../../migrations"
    m, err := migrate.NewWithDatabaseInstance(
        migrationsPath,
    	ms.Cfg.DBName, driver)
    if err != nil {
        return fmt.Errorf("failed to create migrate instance: %w", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("failed to apply migrations: %w", err)
    }

    return nil
}
