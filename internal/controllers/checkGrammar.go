package controllers

import (
	// "encoding/json"
	"log"
	"io"
	"net/http"
	"strconv"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/services"
)

// Check Markdown Grammar
// Input: markdown file
// Output: grammar erros and return nil if none
func CheckGrammar(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	// var note models.Note
	// err := decoder.Decode(&note)
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Malformed")
	}
	note := models.Note{ ID: 1, Markdown: services.MarkdownToHTML(requestBody) }
	w.Write([]byte("Note Markdown: " + string(note.Markdown) + "\nID: " + strconv.Itoa(note.ID) + "\n"))
}
