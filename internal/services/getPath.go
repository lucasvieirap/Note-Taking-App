package services

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
	dir := strings.Split(absPath, "Note-Taking-App")[0] + "Note-Taking-App/"
	return dir
}
