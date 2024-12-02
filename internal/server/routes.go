package server

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/controllers"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"net/http"
)

// TODO: Create dynamic routing for every entry on notes

func RouteHome(mux *http.ServeMux) {
	mux.HandleFunc("GET /", controllers.HandleHomeRouter)
}

func RouteGrammar(mux *http.ServeMux) {
	mux.HandleFunc("GET /grammar", controllers.HandleCheckGrammar)
}

func RouteRemove(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("DELETE /remove", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleRemoveNote(w, r, *storage)
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
