package controllers

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
	"log"
	"net/http"
	"strconv"
)

func HandleListNotes(w http.ResponseWriter, r *http.Request, storage storage.Storage) {
	noteList := storage.GetAll()

	var err error
	for _, note := range noteList {
		note.Markdown, err = util.Decrypt([]byte(note.Markdown), []byte(util.GetEnvValue("AESKEY", "")))

		if err != nil {
			log.Println(err)
		}

		noteId := strconv.Itoa(int(note.ID))
		w.Write([]byte("Note ID: " + noteId + " ===========================================\n"))
		w.Write(util.MarkdownToHTML(note.Markdown))
	}
}

