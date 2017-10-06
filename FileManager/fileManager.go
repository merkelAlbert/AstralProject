//Данный пакет предназначен для удобной работы с файлами
package fileManager

import (
	"AstralProject/Logger"
	"fmt"
	"io/ioutil"
	"os"
)

//Данная функция открывает файл и возвращает его указатель
func Open(s string, flag int, perm os.FileMode) *os.File {
	file, err := os.OpenFile(s, flag, perm)
	if err != nil {
		logger.LogError(err, fmt.Sprintf("Can not open file '%s'", s))
	}
	return file
}

//Данная функция открывает и считывает данные из заданного файла
func Read(s string, flag int, perm os.FileMode) []byte {
	var file = open(s, flag, perm)
	defer file.Close()
	input, err := ioutil.ReadAll(file)
	if err != nil {
		logger.LogError(err, fmt.Sprintf("Can not read file '%s'", s))
	}
	return input
}
