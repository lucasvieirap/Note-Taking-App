package server

import (
	"log"
	"database/sql"
	"net/http"
	
	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
)

type Server struct {
	Address string
	Router *http.ServeMux
	store  storage.Storage
}

func NewServer(address string, router *http.ServeMux, store storage.Storage) *Server {
	return &Server {
		address,
		router,
		store,
	}
}

func (s *Server) Listen() error {
	RouteHome(s.Router)
	RouteGrammar(s.Router)
	RouteList2(s.Router, &s.store)
	RouteSave2(s.Router, &s.store)
	RouteRemove2(s.Router, &s.store)

	return http.ListenAndServe(s.Address, s.Router)
}

func Listen(PORT string, db *sql.DB, mux *http.ServeMux) {
	RouteHome(mux)
	RouteGrammar(mux)
	RouteList(mux, db)
	RouteSave(mux, db)
	RouteRemove(mux, db)

	log.Println("Server Running and Listening at " + PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
