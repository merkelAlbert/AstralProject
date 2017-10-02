package CmdParser

import (
	"flag"
)

const DEFAULTLOGGER = "logger.txt"
const DEFAULTLCONFIG = "person.xml"

func GetPathsFromCommandLine() (string, string) {
	var loggerPath = flag.String("loggerPath", DEFAULTLOGGER, "Setting path for logger file")
	var configPath = flag.String("configPath", DEFAULTLCONFIG, "Setting path for XML-config file")
	flag.Parse()
	return *loggerPath, *configPath
}
