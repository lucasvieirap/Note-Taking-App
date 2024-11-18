package server

import (
	"net/http"

	"github.com/lucasvieirap/Note-Taking-App/internal/controllers"
)

func HomeRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /", controllers.HomeHandler)
}

func GrammarRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /grammar", controllers.CheckGrammar)
}

func SaveRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /save", controllers.SaveNote)
}

func ListRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /list", controllers.ListNotes)
}
