package controllers

import (
	"net/http"
)

// Save markdown file in database
func SaveNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Save\n"))
}
