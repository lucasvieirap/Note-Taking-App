package controllers

import (
	"net/http"
)

// List all markdown files in database
func ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello List\n"))
}
