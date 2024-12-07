package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucasvieirap/Note-Taking-App/internal/models"
)

// Welcome user and Display Usability
func HandleHomeRouter(w http.ResponseWriter, r *http.Request) {
	// TODO: encode to json
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var message models.Message
	message.Message =  
		"\n"+
		"Welcome to my Markdown Note-Taking APP!\n" +
		"Possible Commands Endpoints are:\n" +
		"POST /grammar -> Check your markdown files for grammar errors\n"+
		"POST /save    -> Save your markdown file on the database\n"+
		"GET  /list    -> List markdown files present on the database\n"+
		"\n"

	jsonMessage, err := json.Marshal(message)

	if err != nil {
		log.Println("Error on JSONFY Message")
	}

	w.Write(jsonMessage)
}
