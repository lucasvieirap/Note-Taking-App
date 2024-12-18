package controllers

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"log"
)

func HandleRemoveNote(w http.ResponseWriter, r *http.Request, store storage.Storage, noteId uint) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var message models.Message

	if noteId < 0 {
		log.Println("ID Not a valid number: " + strconv.Itoa(int(noteId)))
		w.Write([]byte("Invalid ID\n"))
	}

	message.Message = "Removed note with id: " + strconv.Itoa(int(noteId)) + "\n"
	jsonMessage, err := json.Marshal(message.Message)
	if err != nil {
		log.Println("Error on MarshalJSON")
	}

	store.Delete(noteId)
	w.Write(jsonMessage)
}
