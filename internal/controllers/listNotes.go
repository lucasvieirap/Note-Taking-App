package controllers

import (
	"net/http"
)

// List all markdown files in database
// And render it as HTML
func ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello List\n"))
}
