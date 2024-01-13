package util

import (
	"io"
	"log"
	"os"
)

var logFile *os.File // パッケージ変数としてファイルポインタを保持

func LoggingSetting(logFilePath string) {
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open " + logFilePath + ": " + err.Error())
	}

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Complete log settings")
}

func CloseLogFile() {
	err := logFile.Close()
	if err != nil {
		panic("cannnot close logfile " + ": " + err.Error())
	}
}
