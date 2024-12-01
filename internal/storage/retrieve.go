package storage

import (
	"log"
	"database/sql"
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
)

func RetrieveAll(db *sql.DB) []models.Note {
	sql := "SELECT * FROM notes";
	rowList, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	defer rowList.Close()

	var noteList []models.Note

	for rowList.Next() {
		note := models.Note{ ID: 0, Markdown: []byte("") }
		var data string
		if err := rowList.Scan(&note.ID, &data); err != nil {
			return nil
		}

		note.Markdown = []byte(data)
		if err != nil {
			log.Panic(err)
		}

		noteList = append(noteList, note)
	}

	if err = rowList.Err(); err != nil {
		return nil
	}

	return noteList
}
