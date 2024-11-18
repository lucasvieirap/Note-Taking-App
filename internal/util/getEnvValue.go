package util

import (
	"log"
	"os"
)

func GetEnvValue(field string, fallback string) string {
	result, ok := os.LookupEnv(field)
	if !ok {
		log.Println("Env variable " + field + " not present")
		return fallback
	}
	return result
}
