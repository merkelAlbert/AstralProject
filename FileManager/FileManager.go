package FileManager

import (
	"AstralProject/Logger"
	"fmt"
	"io/ioutil"
	"os"
)

//Open returns file descriptor with path s, flag to open file, and permission
func open(s string, flag int, perm os.FileMode) *os.File {
	file, err := os.OpenFile(s, flag, perm)
	if err != nil {
		Logger.LogError(err, fmt.Sprintf("Can not openfile '%s'", s))
	}
	openedFile = file
	return file
}

//Read opens and reads and returns information from file with path s, flag to open file, and permission
func Read(s string, flag int, perm os.FileMode) []byte {
	var file = open(s, flag, perm)
	defer file.Close()
	input, err := ioutil.ReadAll(file)
	if err != nil {
		Logger.LogError(err, fmt.Sprintf("Can not read file '%s'", s))
	}
	return input
}
