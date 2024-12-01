package controllers

import (
	"net/http"
)

// Check Markdown Grammar
// Input: markdown file
// Output: grammar erros and return nil if none
func HandleCheckGrammar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Grammar"))
}
