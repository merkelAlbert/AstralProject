package main

import (
	"AstralProject/CmdParser"
	"AstralProject/ConfigTypes"
	"AstralProject/FileManager"
	//"AstralProject/github.com/lib/pq"
	//"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func main() {
	var configPath string
	var loggerPath string
	var config []byte

	loggerPath, configPath = cmdParser.GetPathsFromCommandLine()

	logger, err := os.OpenFile(loggerPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(logger)
		log.Printf("log file %s is succesfully opened", loggerPath)
	}

	config = fileManager.Read(configPath, os.O_RDONLY, 0444)

	var data configTypes.Data
	err = xml.Unmarshal(config, &data)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("config %s is succesful unmarshalled", configPath)

	}

	for _, v := range data.PersonList {
		fmt.Println(v)
	}

}
