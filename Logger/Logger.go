package Logger

import (
	"fmt"
	"log"
	"os"
)

//SetLogger sets file(path) where logs will be written
func SetLogger(path string) {
	var logger *os.File
	logger, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	LogError(err, fmt.Sprintf("Can not create file '%s'", path))
	log.SetOutput(logger)
}

//LogError logs error and description to established log file
func LogError(err error, description string) {
	if err != nil {
		log.Println("Error: ", err.Error(), ". ", description)
	}
}
