package controllers

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
)

func HandleRemoveNote(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("No id on DELETE request")
	}

	deleteNoteId, err := strconv.Atoi(string(requestBody))
	if err != nil || deleteNoteId < 0{
		log.Println("ID Not a valid number: " + string(requestBody))
		w.Write([]byte("Invalid ID\n"))
	}

	store.Delete(uint(deleteNoteId))
}
