package Logger

import (
	"fmt"
	"log"
	"os"
)

var logger *os.File

func SetLogger(path string) {
	logger, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	LogError(err, fmt.Sprintf("Can not create file '%s'", path))
	log.SetOutput(logger)
}

func LogError(err error, description string) {
	if err != nil {
		log.Println("Error: ", err.Error(), ". ", description)
	}
}

func LogText(text string) {
	log.Println(text)
}
