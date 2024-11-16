package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/lucasvieirap/Note-Taking-App/internal/routes"
	"github.com/lucasvieirap/Note-Taking-App/internal/env"
	"github.com/lucasvieirap/Note-Taking-App/internal/services"
)

func main() {
	mux := http.NewServeMux()

	routes.HomeRouter(mux)
	routes.ListRouter(mux)
	routes.SaveRouter(mux)
	routes.GrammarRouter(mux)

	dir := services.GetPath()

	godotenv.Load(dir + ".env")
	PORT := env.GetEnvValue("PORT", ":8000")

	log.Println("Server Running and Listening at " + PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
