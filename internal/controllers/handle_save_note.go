package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
)

func HandleSaveNote(w http.ResponseWriter, r *http.Request, storage storage.Storage) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var message models.Message

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error on Request")
	}

	note := models.Note{ ID: 0, Markdown: requestBody }
	if err != nil {
		log.Println("Error on JSON Parsing.")
	}

	encryptedBody, err := util.Encrypt(note.Markdown, []byte(util.GetEnvValue("AESKEY", "")))
	if err != nil {
		log.Panic(err)
	}

	encryptedNote := models.Note{ ID: note.ID, Markdown: encryptedBody }
	
	encryptedNote.ID, err = storage.Create(encryptedNote)
	if err != nil {
		log.Println("Error while inserting in DB: ", err)
	}

	log.Println("EncryptedNote: " + string(encryptedNote.Markdown))
	log.Println("Id: " + strconv.Itoa(int(encryptedNote.ID)))

	message.Message = "Saved note with id: " + strconv.Itoa(int(encryptedNote.ID)) + "\n" 

	jsonMessage, err := json.Marshal(message)

	if err != nil {
		log.Println("Error on JSONFY Message")
	}
	
	w.Write(jsonMessage)
}
