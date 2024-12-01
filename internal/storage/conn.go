package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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

func OpenDBConn(cfg DBConnConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetConfigString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("DB Server Running: " + cfg.Host + ":" + cfg.Port)

	return db, err
}
