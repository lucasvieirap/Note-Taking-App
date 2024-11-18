package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/lucasvieirap/Note-Taking-App/internal/database"
	"github.com/lucasvieirap/Note-Taking-App/internal/server"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
)

func main() {
	mux := http.NewServeMux()

	dir := util.GetPath()
	godotenv.Load(dir + "/.env")
	PORT := util.GetEnvValue("PORT", ":8000")

	db, err := database.OpenConn()
	if err != nil {
		log.Panic(err)
	}

	// Add to scripts
	db.Exec("CREATE DATABASE note_taking_app;")
	db.Exec("CREATE notes (id BIGSERIAL PRIMARY KEY NOT NULL, data VARCHAR(10000));")

	log.Println("DB Open Connectionss: " + strconv.Itoa(db.Stats().OpenConnections))

	server.Listen(PORT, mux)
}
