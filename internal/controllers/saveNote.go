package controllers

import (
	// "encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/database"
)

// Save markdown file in database
func SaveNote(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error on Request")
	}

	note := models.Note{ ID: 1, Markdown: requestBody, }
	// jsonote, err := json.Marshal(note)
	if err != nil {
		log.Println("Error on JSON Parsing.")
	}
	
	database.Insert(note)
	w.Write([]byte("Saved note with id: " + strconv.Itoa(note.ID)))
}
