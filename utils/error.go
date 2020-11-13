package utils

import (
	"log"
	"os"
)

type PrintDefaults func()

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func InvalidAction(message string) {
	log.Printf(message)
	os.Exit(1)
}
