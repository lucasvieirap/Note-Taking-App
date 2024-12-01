package storage

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"database/sql"
)

func Insert(n models.Note, db *sql.DB) (id uint, err error) {
	sql := `INSERT INTO notes (data) VALUES ($1) RETURNING id`
	db.QueryRow(sql, n.Markdown).Scan(&id)

	return id, err
}
