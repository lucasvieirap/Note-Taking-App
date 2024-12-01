package controllers

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"database/sql"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
)

func HandleRemoveNote(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("No id on DELETE request")
	}

	storage.RemoveFromTableUsingId(db, requestBody)
}

func HandleRemoveNote2(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("No id on DELETE request")
	}

	deleteNoteId, err := strconv.Atoi(string(requestBody))
	if err != nil || deleteNoteId < 0{
		log.Println("ID Not a valid number.")
	}

	store.Delete(uint(deleteNoteId))
}
