package controllers

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
	"log"
	"net/http"
	"encoding/json"
	// "strconv"
)

type noteListModel struct {
	NoteList []models.Note `json:"noteList"`
}

func HandleListNotes(w http.ResponseWriter, r *http.Request, storage storage.Storage) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	var noteList noteListModel
	noteList.NoteList = storage.GetAll()

	var decryptedNoteList noteListModel

	var err error
	for _, note := range noteList.NoteList {
		note.Markdown, err = util.Decrypt(note.Markdown, []byte(util.GetEnvValue("AESKEY", "")))

		if err != nil {
			log.Println(err)
		}

		decryptedNoteList.NoteList = append(decryptedNoteList.NoteList, note)
	}

	jsonMessage, err := json.Marshal(decryptedNoteList)

	if err != nil {
		log.Println("Error on JSONFY Message")
	}

	w.Write(jsonMessage)
}

