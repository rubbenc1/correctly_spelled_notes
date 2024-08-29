package main

import (
	"net/http"
	"note-service/config"
	"note-service/internal/handlers"
	"note-service/internal/logging"
	"note-service/internal/storage"
	"note-service/internal/middleware"
	"log/slog"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := logging.InitializeLogger(cfg.Logging.Level)
	logger.Info("Starting the application...", slog.String("port", cfg.Server.Port))

	// Initialize database connection
	db, err := storage.InitDB(&cfg.Postgres)
	if err != nil {
		logger.Error("failed to connect to database", slog.String("error", err.Error()))
		return
	}
	defer db.Close()

	// Run database migrations
	if err := storage.NewMigrationService(db, cfg.Postgres).Run(); err != nil {
		logger.Error("failed to run migrations", slog.String("error", err.Error()))
		return
	}

	// Initialize the note storage service
	noteStorage := &storage.NoteStorage{DB: db}

	http.HandleFunc("/notes", middleware.AuthMiddleware(logger, handlers.AddNoteHandler(logger, noteStorage)))          // Add a new note
	http.HandleFunc("/notes/", middleware.AuthMiddleware(logger, handlers.ListUserNotesHandler(logger, noteStorage))) // List notes for a user

	// Start the HTTP server
	logger.Info("Server is starting...", slog.String("port", cfg.Server.Port))
	if err := http.ListenAndServe(":"+cfg.Server.Port, nil); err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
	}
}
