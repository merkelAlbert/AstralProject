package main

import (
	"AstralProject/CmdParser"
	"AstralProject/ConfigTypes"
	"AstralProject/FileManager"
	"AstralProject/Logger"
	"encoding/xml"
	"fmt"
	"os"
)

var configPath string
var loggerPath string
var config []byte

func main() {
	loggerPath, configPath = CmdParser.GetPathsFromCommandLine()
	Logger.SetLogger(loggerPath)

	config = FileManager.Read(configPath, os.O_RDONLY, 0444)
	var data ConfigTypes.Data

	xml.Unmarshal(config, &data)
	for _, v := range data.PersonList {
		fmt.Println(v)
	}
}
