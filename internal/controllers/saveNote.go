package controllers

import (
	// "encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/database"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
)

// Save markdown file in database
func SaveNote(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error on Request")
	}

	note := models.Note{ ID: 1, Markdown: requestBody, }
	if err != nil {
		log.Println("Error on JSON Parsing.")
	}

	encryptedBody, err := util.Encrypt(note.Markdown, []byte(util.GetEnvValue("AESKEY", "")))
	encryptedNote := models.Note{ ID: note.ID, Markdown: encryptedBody }
	
	database.Insert(encryptedNote)
	w.Write([]byte("Saved note with id: " + strconv.Itoa(note.ID)))
}
