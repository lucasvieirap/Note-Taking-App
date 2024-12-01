package storage

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
	"database/sql"
	"log"
	"os"
	"strconv"
)

type PostgresStorage struct {
	DBConn *sql.DB
}

func NewPostgresStorage(dbConfig DBConnConfig) *PostgresStorage {
	dbConn, err := OpenDBConn(dbConfig)

	initQueryFile, err := os.Open(util.GetPath() + "/scripts/init.sql")
	if err != nil {
		log.Fatal(err)
	}

	initQuery := make([]byte, 100)
	initQueryFile.Read(initQuery)

	dbConn.Exec(string(initQuery))

	log.Println("DB Open Connectionss: " + strconv.Itoa(dbConn.Stats().OpenConnections))

	if err != nil {
		log.Fatal("Error Connecting to DB")
	}

	return &PostgresStorage {
		dbConn,
	}
}

func (p *PostgresStorage) Get() *models.Note {
	return &models.Note{
		ID: 0,
		Markdown: []byte(""),
	}
}

func (p *PostgresStorage) GetAll() []models.Note {
	sql := "SELECT * FROM notes";
	rowList, err := p.DBConn.Query(sql)
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

func (p *PostgresStorage) Create(note models.Note) (id uint, err error) {
	sql := `INSERT INTO notes (data) VALUES ($1) RETURNING id`
	p.DBConn.QueryRow(sql, note.Markdown).Scan(&id)

	return id, err
}

func (p *PostgresStorage) Delete(id uint) {
	sql := "DELETE FROM notes WHERE id = ($1)"
	p.DBConn.Exec(sql, id)
}
