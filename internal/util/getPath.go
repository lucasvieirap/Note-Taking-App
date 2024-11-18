package util 

import (
	"log"
	"path/filepath"
	"strings"
)

func GetPath() string {
	absPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Error on getting Path")
	}
	dir := filepath.Join(strings.Split(absPath, "Note-Taking-App")[0], "Note-Taking-App")
	log.Println("project path: " + dir)
	return dir
}
