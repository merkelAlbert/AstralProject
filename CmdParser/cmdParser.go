//Данный пакет предназначен для считывания аргументов из командной строки
package cmdParser

import (
	"AstralProject/Constants"
	"flag"
)

//Данная функция получает из командной строки путь для логирования
//и путь конфига и возвращает их в этой же последовательнояти
func GetPathsFromCommandLine() (string, string) {
	var loggerPath = flag.String("loggerPath", constants.DEFAULTLOGGER, "Setting path for logger file")
	var configPath = flag.String("configPath", constants.DEFAULTCONFIG, "Setting path for XML-config file")
	flag.Parse()
	return *loggerPath, *configPath
}
