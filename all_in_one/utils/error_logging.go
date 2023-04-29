package utils

import (
	"log"
	"os"
)

func Logging(message string, err error) {
	f, er := os.OpenFile("errors.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if er != nil {
		log.Fatalf("failed to open log file: %v", er)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Printf("%s: %v", message, err)
}
