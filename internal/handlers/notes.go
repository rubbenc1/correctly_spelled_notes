package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"note-service/internal/services"
	"note-service/internal/storage"
	"time"
)

func AddNoteHandler(logger *slog.Logger, noteStorage *storage.NoteStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("userID").(int)
		if !ok {
			http.Error(w, "User not authorized", http.StatusUnauthorized)
			return
		}

		var req struct {
			Text string `json:"text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate the note's text using Yandex Speller
		spellingResults, err := services.CheckSpelling([]string{req.Text})
		if err != nil {
			logger.Error("failed to validate spelling", slog.String("error", err.Error()))
			http.Error(w, "Failed to validate spelling", http.StatusInternalServerError)
			return
		}

		if len(spellingResults) > 0 && len(spellingResults[0]) > 0 {
			corrections := formatSpellerResults(spellingResults[0])
			logger.Info("spelling corrections", slog.String("corrections", corrections))
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":       "Spelling mistakes found",
				"corrections": corrections,
			})
			return
		}

		createdAt := time.Now()
		if err := noteStorage.SaveNote(userID, req.Text, createdAt); err != nil {
			logger.Error("failed to save note", slog.String("error", err.Error()))
			http.Error(w, "Failed to save note", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Note added successfully",
		})
	}
}

// formatSpellerResults formats spelling suggestions
func formatSpellerResults(results []services.SpellerError) string {
	var formatted string
	for _, result := range results {
		if len(result.Suggestions) > 0 {
			formatted += fmt.Sprintf("'%s' -> '%s' ", result.Word, result.Suggestions[0])
		}
	}
	return formatted
}