package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lucasvieirap/Note-Taking-App/internal/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.HomeRouter(mux)
	routes.ListRouter(mux)
	routes.SaveRouter(mux)
	routes.GrammarRouter(mux)

	absPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Error on loading envirement file")
	}
	dir := strings.Split(absPath, "Note-Taking-App")[0] + "Note-Taking-App"

	godotenv.Load(dir + "/internal/env/.env")
	PORT := os.Getenv("PORT")

	log.Println("Server Running and Listening at " + PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
