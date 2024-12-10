package controllers

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/models"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
)

type noteListModel struct {
	NoteList []models.Note `json:"noteList"`
}

func (nl noteListModel) MarshalJSON() ([]byte, error) {
	jsonData := `{"noteList":[`

	for index, note := range nl.NoteList {
		markdownReplaced := strings.Replace(strings.Replace(string(note.Markdown), "\n", "\\n", -1), "\"", "'", -1)
		jsonData += `{"id":"` + strconv.Itoa(int(note.ID)) + `", ` + `"note":"` + markdownReplaced + `"}`
		if index < len(nl.NoteList) -1 {
			jsonData += ", "
		}
	}

	jsonData +=  `]}`

	return []byte(jsonData), nil
}

func HandleListNotes(w http.ResponseWriter, r *http.Request, storage storage.Storage) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	var noteList noteListModel
	noteList.NoteList = storage.GetAll()

	var decryptedNoteList noteListModel

	var err error
	for _, note := range noteList.NoteList {
		note.Markdown, err = util.Decrypt(note.Markdown, []byte(util.GetEnvValue("AESKEY", "")))

		if err != nil {
			log.Println(err)
		}

		decryptedNoteList.NoteList = append(decryptedNoteList.NoteList, note)
	}

	jsonMessage, err := json.Marshal(decryptedNoteList)
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte(string(jsonMessage) + "\n"))
}

