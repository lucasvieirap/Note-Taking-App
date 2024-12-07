package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
)

func HandleRemoveNote(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var message models.Message

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("No id on DELETE request")
	}

	deleteNoteId, err := strconv.Atoi(string(requestBody))
	if err != nil || deleteNoteId < 0{
		log.Println("ID Not a valid number: " + string(requestBody))
		w.Write([]byte("Invalid ID\n"))
	}

	message.Message = "Removed note with id: " + strconv.Itoa(deleteNoteId)
	jsonMessage, err := json.Marshal(message.Message)
	if err != nil {
		log.Println("Error on MarshalJSON")
	}

	store.Delete(uint(deleteNoteId))
	w.Write(jsonMessage)
}
