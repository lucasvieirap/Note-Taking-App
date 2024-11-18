package server

import (
	"log"
	"net/http"
)

func Listen(PORT string, mux *http.ServeMux) {
	HomeRouter(mux)
	ListRouter(mux)
	SaveRouter(mux)
	GrammarRouter(mux)

	log.Println("Server Running and Listening at " + PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
