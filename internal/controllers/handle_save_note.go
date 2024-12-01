package controllers

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"database/sql"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
)

func HandleSaveNote(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error on Request")
	}

	note := models.Note{ ID: 0, Markdown: requestBody, }
	if err != nil {
		log.Println("Error on JSON Parsing.")
	}

	encryptedBody, err := util.Encrypt(note.Markdown, []byte(util.GetEnvValue("AESKEY", "")))
	if err != nil {
		log.Panic(err)
	}

	encryptedNote := models.Note{ ID: note.ID, Markdown: encryptedBody }
	
	note.ID, err = storage.Insert(encryptedNote, db)
	if err != nil {
		log.Println("Error while inserting in DB")
	}
	
	w.Write([]byte("Saved note with id: " + strconv.Itoa(int(encryptedNote.ID)) + "\n"))
}

func HandleSaveNote2(w http.ResponseWriter, r *http.Request, storage storage.Storage) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error on Request")
	}

	note := models.Note{ ID: 0, Markdown: requestBody, }
	if err != nil {
		log.Println("Error on JSON Parsing.")
	}

	encryptedBody, err := util.Encrypt(note.Markdown, []byte(util.GetEnvValue("AESKEY", "")))
	if err != nil {
		log.Panic(err)
	}

	encryptedNote := models.Note{ ID: note.ID, Markdown: encryptedBody }
	
	note.ID, err = storage.Create(encryptedNote)
	if err != nil {
		log.Println("Error while inserting in DB")
	}
	
	w.Write([]byte("Saved note with id: " + strconv.Itoa(int(encryptedNote.ID)) + "\n"))
}
