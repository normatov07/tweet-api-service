package utils

import (
	"log"
	"os"
)

var LogFile *os.File

func LoadLogs() {
	var err error
	LogFile, err = os.OpenFile("logs/system.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	log.SetOutput(LogFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}
