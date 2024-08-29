package logging

import (
    "os"
    "log/slog"
)

// InitializeLogger sets up a structured logger based on the log level provided.
func InitializeLogger(logLevel string) *slog.Logger {
    var logger *slog.Logger
    switch logLevel {
    case "debug":
        logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case "info":
        logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
    case "error":
        logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
    default:
        logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
    }

    return logger
}
