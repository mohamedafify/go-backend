package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func getLoggerOutput() *io.WriteCloser {
	var logFile io.WriteCloser
	gin.DisableConsoleColor()
	// DD-MM-YYYY_HH24:MI:SS.log
	logFileName := fmt.Sprintf("%v.log", time.Now().Local().Format("02-01-2006_15:04:05"))
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to create logfile: %v", err.Error()))
		log.Println("Using Stdout to log")
		logFile = os.Stdout
	} else {
		logFile = file
	}
	return &logFile
}

func SetupLoggers() *io.WriteCloser {
	logFile := getLoggerOutput()
	gin.DefaultWriter = *logFile
	log.SetOutput(*logFile)
	return logFile
}
