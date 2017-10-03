package FileManager

import (
	"AstralProject/Logger"
	"fmt"
	"io/ioutil"
	"os"
)

func Open(s string, flag int, perm os.FileMode) *os.File {
	file, err := os.OpenFile(s, flag, perm)
	Logger.LogError(err, fmt.Sprintf("Can not openfile '%s'", s))
	return file
}

func Read(s string) []byte {
	var file = Open(s, os.O_RDONLY, 0444)
	defer file.Close()
	input, err := ioutil.ReadAll(file)
	Logger.LogError(err, fmt.Sprintf("Can not read file '%s'", s))
	return input
}
