package server

import (
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
	RouteHomePage(s.Router)
	RouteGrammar(s.Router)
	RouteList(s.Router, &s.store)
	RouteSave(s.Router, &s.store)
	RouteRemove(s.Router, &s.store)

	return http.ListenAndServe(s.Address, s.Router)
}
