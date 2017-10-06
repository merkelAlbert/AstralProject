//Данный пакет логирует информацию о работе программы.
//Он содержит в себе две функции. Первая для установки файла для логирования,
//а вторая выводит информацию в заданный файл
package logger

import (
	"AstralProject/Constants"
	"fmt"
	"log"
	"os"
)

//Глобальная переменная, хранящая указать на файл для логирования
var logger *os.File

//Данная функция устанавливает выходной файл для логирования
func SetLogger(path string) {
	logger, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err, fmt.Sprintf("Can not create file '%s'", path))
		return
	}
	log.SetOutput(logger)
}

//Данная функция пишет ошибку и ее описание в установленный файл или же в консоль,если файл не был установлен
func LogError(err error, description string) {
	if logger == nil {
		SetLogger(constants.DEFAULTLOGGER)
	}
	if err != nil {
		log.Println("Error: ", err.Error(), ". ", description)
	}
}
