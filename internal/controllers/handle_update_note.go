package controllers

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"io"
)

func HandleUpdateNote(w http.ResponseWriter, r *http.Request, storage storage.Storage, noteId uint) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error on Request")
	}

	note := storage.Get(noteId)
	if err != nil {
		log.Println("Error on JSON Parsing.")
	}

	note.Markdown = requestBody

	encryptedBody, err := util.Encrypt(note.Markdown, []byte(util.GetEnvValue("AESKEY", "")))
	if err != nil {
		log.Panic(err)
	}

	encryptedNote := models.Note{ ID: note.ID, Markdown: encryptedBody }

	var message models.Message

	if noteId < 0 {
		log.Println("ID Not a valid number: " + strconv.Itoa(int(noteId)))
		w.Write([]byte("Invalid ID\n"))
	}

	storage.Update(noteId, encryptedNote)

	message.Message = "Updated Successful - Note With id: " + strconv.Itoa(int(noteId)) + "\n"
	jsonMessage, err := json.Marshal(message.Message)
	if err != nil {
		log.Println("Error on MarshalJSON")
	}

	w.Write(jsonMessage)
}
