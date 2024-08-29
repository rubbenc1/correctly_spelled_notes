package storage

import (
	"database/sql"
	"fmt"
	"note-service/internal/models"
	"time"
)

type NoteStorage struct {
	DB *sql.DB
}

func (ns *NoteStorage) SaveNote(userID int, text string, createdAt time.Time) error {
	query := `INSERT INTO notes (user_id, text, created_at) VALUES ($1, $2, $3)`
	_, err := ns.DB.Exec(query, userID, text, createdAt)
	if err != nil {
		return fmt.Errorf("failed to save note: %w", err)
	}
	return nil
}

func (ns *NoteStorage) ListNotes(userID int) ([]models.Note, error){
	query := `SELECT id, user_id, text, created_at FROM notes WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := ns.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list note: %w", err)
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.UserID, &note.Text, &note.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return notes, nil
}