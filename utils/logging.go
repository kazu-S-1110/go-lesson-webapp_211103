package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// logFileを読み書き可能で且つなければ作成して追記形式で読み出す。
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
