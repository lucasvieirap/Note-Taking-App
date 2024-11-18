package database

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
)

func Insert(n models.Note) (id int, err error) {
	db, err := OpenConn()
	if err != nil {
		return
	}

	defer db.Close()

	sql := `INSERT INTO notes (data) VALUES ($1) RETURNING id`
	err = db.QueryRow(sql, n.Markdown).Scan(&id)

	return id, err
}
