package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/lucasvieirap/Note-Taking-App/internal/util"
)

type DBConnConfig struct {
	Host	 string
	Port	 string
	User	 string
	Password string
	DBname	 string
	SSLmode  string
}

func (cfg DBConnConfig) GetConfigString() string {
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
						cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname, cfg.SSLmode)
	return str
}

func OpenConn() (*sql.DB, error) {
	passwd := util.GetEnvValue("DATABASE_PASSWD", " ")

	cfg := DBConnConfig{
		Host: "localhost",
		Port: "5432",
		User: "postgres",
		Password: passwd,
		DBname: "note_taking_app",
		SSLmode: "disable",
	}

	db, err := sql.Open("postgres", cfg.GetConfigString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	return db, err
}
