package main

import (
	"AstralProject/CmdParser"
	"AstralProject/ConfigTypes"
	"AstralProject/FileManager"
	"AstralProject/Logger"
	"AstralProject/github.com/lib/pq"
	"database/sql"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	var configPath string
	var loggerPath string
	var config []byte
	loggerPath, configPath = cmdParser.GetPathsFromCommandLine()
	logger.SetLogger(loggerPath)
	config = fileManager.Read(configPath, os.O_RDONLY, 0444)
	var data configTypes.Data
	xml.Unmarshal(config, &data)
	for _, v := range data.PersonList {
		fmt.Println(v)
	}
}
