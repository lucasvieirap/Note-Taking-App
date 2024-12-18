package server

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/controllers"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"net/http"
	"strings"
	"strconv"
)

func RouteHome(mux *http.ServeMux) {
	mux.HandleFunc("GET /", controllers.HandleHomeRouter)
}

func RouteNotePage(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("GET /note/", func(w http.ResponseWriter, r *http.Request) {
		var page []string
		page = strings.Split(r.URL.Path, "/")
		noteID := page[len(page)-1]
		controllers.HandleNotePageRouter(w, r, *storage, noteID)
	})
}

func RouteUpdateNote(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("PUT /note/", func(w http.ResponseWriter, r *http.Request) {
		var page []string
		page = strings.Split(r.URL.Path, "/")
		noteID := page[len(page)-1]
		intNoteId, _ := strconv.Atoi(noteID)
		controllers.HandleUpdateNote(w, r, *storage, uint(intNoteId))
	})
}

func RouteGrammar(mux *http.ServeMux) {
	mux.HandleFunc("GET /grammar", controllers.HandleCheckGrammar)
}

func RouteRemove(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("DELETE /note/", func(w http.ResponseWriter, r *http.Request) {
		var page []string
		page = strings.Split(r.URL.Path, "/")
		noteID := page[len(page)-1]
		intNoteId, _ := strconv.Atoi(noteID)
		controllers.HandleRemoveNote(w, r, *storage, uint(intNoteId))
	})
}

func RouteList(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("GET /list", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleListNotes(w, r, *storage)
	})
}

func RouteSave(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("POST /save", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleSaveNote(w, r, *storage)
	})
}
