package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/lucasvieirap/Note-Taking-App/internal/storage"
	"github.com/lucasvieirap/Note-Taking-App/internal/server"
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
)

func main() {
	mux := http.NewServeMux()

	dir := util.GetPath()
	godotenv.Load(dir + "/.env")
	PORT := util.GetEnvValue("PORT", ":8000")

	passwd := util.GetEnvValue("DATABASE_PASSWD", " ")
	cfg := storage.DBConnConfig{
		Host: "localhost",
		Port: "5432",
		User: "postgres",
		Password: passwd,
		DBname: "note_taking_app",
		SSLmode: "disable",
	}

	// db, err := storage.OpenDBConn(cfg)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// initQueryFile, err := os.Open(util.GetPath() + "/scripts/init.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// initQuery := make([]byte, 100)
	// initQueryFile.Read(initQuery)

	// db.Exec(string(initQuery))

	// log.Println("DB Open Connectionss: " + strconv.Itoa(db.Stats().OpenConnections))

	// server.Listen(PORT, db, mux)

	postgresStorage := storage.NewPostgresStorage(cfg)
	serverInstance := server.NewServer(PORT, mux, postgresStorage)

	if err := serverInstance.Listen(); err != nil {
		log.Fatal(err)
	}
}
