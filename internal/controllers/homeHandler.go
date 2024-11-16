package controllers

import (
	"net/http"
)

// Welcome user and Display Usability
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	message := 
		"\nWelcome to my Markdown Note-Taking APP!\n" +
		"Possible Commands Endpoints are:\n" +
		"POST /grammar -> Check your markdown files for grammar errors\n"+
		"POST /save    -> Save your markdown file on the database\n"+
		"GET /list     -> List markdown files present on the database\n"
	w.Write([]byte(message + "\n"))
}