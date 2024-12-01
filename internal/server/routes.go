package server

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/controllers"
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"net/http"
	"database/sql"
)

func RouteHome(mux *http.ServeMux) {
	mux.HandleFunc("GET /", controllers.HandleHomeRouter)
}

func RouteGrammar(mux *http.ServeMux) {
	mux.HandleFunc("GET /grammar", controllers.HandleCheckGrammar)
}

func RouteRemove(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("GET /remove", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleRemoveNote(w, r, db)
	})
}

func RouteList(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("GET /list", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleListNotes(w, r, db)
	})
}

func RouteSave(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("POST /save", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleSaveNote(w, r, db)
	})
}

func RouteRemove2(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("GET /remove", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleRemoveNote2(w, r, *storage)
	})
}

func RouteList2(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("GET /list", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleListNotes2(w, r, *storage)
	})
}

func RouteSave2(mux *http.ServeMux, storage *storage.Storage) {
	mux.HandleFunc("POST /save", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleSaveNote2(w, r, *storage)
	})
}
