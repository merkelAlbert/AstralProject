//Данный пакет предназначен для удобной работы с файлами
package fileManager

import (
	"io/ioutil"
	"log"
	"os"
)

//Данная функция открывает и считывает данные из заданного файла//ReadFile
func Read(s string, flag int, perm os.FileMode) []byte {
	file, err := os.OpenFile(s, flag, perm)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("file %s is succefully opened", s)
	}
	defer file.Close()
	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("file %s is succesully readed", s)
	}
	return input
}
