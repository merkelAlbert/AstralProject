package main

import (
	"AstralProject/CmdParser"
	"AstralProject/ConfigTypes"
	"AstralProject/ConfigValidator"
	"AstralProject/FileManager"
	"AstralProject/Logger"
	"encoding/xml"
	"fmt"
)

var configPath string
var loggerPath string
var config []byte

func main() {
	loggerPath, configPath = CmdParser.GetPathsFromCommandLine()
	Logger.SetLogger(loggerPath)

	config = FileManager.Read(configPath)
	var data ConfigTypes.Data

	xml.Unmarshal(config, &data)

	for _, v := range data.PersonList {
		fmt.Println(v)
		fmt.Println(ConfigValidator.ValidateAge(v.Age))
	}
}
