package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"note-service/internal/storage"
	"strconv"
	"strings"
)

func ListUserNotesHandler(logger *slog.Logger, noteStorage *storage.NoteStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("userID").(int)
		if !ok {
			http.Error(w, "User not authorized", http.StatusUnauthorized)
			return
		}

		userIDStr := strings.TrimPrefix(r.URL.Path, "/notes/")
		if _, err := strconv.Atoi(userIDStr); err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		notes, err := noteStorage.ListNotes(userID)
		if err != nil {
			logger.Error("failed to retrieve notes", slog.String("error", err.Error()))
			http.Error(w, "Failed to retrieve notes", http.StatusInternalServerError)
			return
		}

		if len(notes) == 0 {
			http.Error(w, "No notes found for user", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(notes); err != nil {
			logger.Error("failed to encode notes", slog.String("error", err.Error()))
			http.Error(w, "Failed to return notes", http.StatusInternalServerError)
		}
	}
}
