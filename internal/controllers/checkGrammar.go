package controllers

import (
	"net/http"
	"strconv"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
)

// Check Markdown Grammar
// Input: markdown file
// Output: grammar erros and return nil if none
func CheckGrammar(w http.ResponseWriter, r *http.Request) {
	note := models.Note{ ID: 1, Markdown: "<h1>New Note</h1>" }
	w.Write([]byte("Note Markdown: " + note.Markdown + "\nID: " + strconv.Itoa(note.ID) + "\n"))
}
