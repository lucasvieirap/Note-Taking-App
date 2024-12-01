package util 

import (
	"log"
	"path/filepath"
	"strings"
)

func GetPath() string {
	absolutePath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Error on getting Path")
	}

	currentDirectory := filepath.Join(strings.Split(absolutePath, "Note-Taking-App")[0], "Note-Taking-App")
	return currentDirectory
}
